package cli

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ory/graceful"
	slogchi "github.com/samber/slog-chi"
	"github.com/spf13/cobra"
	"github.com/swaggest/swgui/v5emb"
	"github.com/ugent-library/httpx/render"
	"github.com/ugent-library/projects-service/api/v1"
	"github.com/ugent-library/projects-service/repositories"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

type apiSecurityHandler struct {
	APIKey string
}

func (s *apiSecurityHandler) HandleApiKey(ctx context.Context, operationName string, t api.ApiKey) (context.Context, error) {
	if t.APIKey == s.APIKey {
		return ctx, nil
	}
	return ctx, errors.New("unauthorized")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// setup repo
		repo, err := repositories.New(repositories.Config{
			Conn: config.Repo.Conn,
		})
		if err != nil {
			return err
		}

		// setup searcher
		// searcher, err := es6.NewSearcher(es6.Config{
		// 	Conn:  strings.Split(config.Search.Conn, ","),
		// 	Index: config.Search.Index,
		// })
		// if err != nil {
		// 	return err
		// }

		// setup api
		log.Printf("%s", config.APIKey)
		apiServer, err := api.NewServer(api.NewService(repo), &apiSecurityHandler{APIKey: config.APIKey})
		if err != nil {
			return err
		}

		// setup mux
		mux := chi.NewMux()
		mux.Use(middleware.RequestID)
		if config.Env != "local" {
			mux.Use(middleware.RealIP)
		}
		mux.Use(slogchi.NewWithConfig(logger, slogchi.Config{
			WithRequestID: true,
		}))
		mux.Use(middleware.Recoverer)

		// mount health and info
		mux.Get("/health", health.NewHandler(health.NewChecker())) // TODO add checkers
		mux.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, http.StatusOK, &struct {
				Branch string `json:"branch,omitempty"`
				Commit string `json:"commit,omitempty"`
				Image  string `json:"image,omitempty"`
			}{
				Branch: version.Branch,
				Commit: version.Commit,
				Image:  version.Image,
			})
		})

		// mount api
		mux.Mount("/api/v1", http.StripPrefix("/api/v1", apiServer))
		mux.Get("/api/v1/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "api/v1/openapi.yaml")
		})
		mux.Mount("/api/v1/docs", v5emb.New(
			"Projects service",
			"/api/v1/openapi.yaml",
			"/api/v1/docs",
		))

		// start server
		server := graceful.WithDefaults(&http.Server{
			Addr:         config.Addr(),
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
		logger.Infof("starting server at %s", config.Addr())
		if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
			return err
		}
		logger.Info("gracefully stopped server")

		return nil
	},
}

package cli

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/graceful"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	slogchi "github.com/samber/slog-chi"
	"github.com/spf13/cobra"
	"github.com/swaggest/swgui/v5emb"
	"github.com/ugent-library/httpx/render"
	"github.com/ugent-library/projects-service/api/v1"
	"github.com/ugent-library/projects-service/indexes"
	"github.com/ugent-library/projects-service/jobs"
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
		ctx := context.Background()

		// setup repo
		pool, err := pgxpool.New(ctx, config.Repo.Conn)
		if err != nil {
			return err
		}

		repo, err := repositories.New(repositories.RepoConfig{
			Conn: pool,
		})
		if err != nil {
			return err
		}

		// Setup index
		index, err := indexes.NewIndex(indexes.IndexConfig{
			Conn:      config.Index.Conn,
			Name:      config.Index.Name,
			Retention: config.Index.Retention,
			Logger:    logger,
		})
		if err != nil {
			return err
		}

		// start job server
		riverWorkers := river.NewWorkers()
		river.AddWorker(riverWorkers, jobs.NewReindexProjectsWorker(repo, index))
		riverClient, err := river.NewClient(riverpgxv5.New(pool), &river.Config{
			Logger:  logger,
			Workers: riverWorkers,
			Queues: map[string]river.QueueConfig{
				river.QueueDefault: {MaxWorkers: 100},
			},
			PeriodicJobs: []*river.PeriodicJob{
				river.NewPeriodicJob(
					river.PeriodicInterval(30*time.Minute),
					func() (river.JobArgs, *river.InsertOpts) {
						return jobs.ReindexProjectsArgs{}, nil
					},
					&river.PeriodicJobOpts{RunOnStart: true},
				),
			},
		})
		if err != nil {
			return err
		}
		if err := riverClient.Start(ctx); err != nil {
			return err
		}
		defer riverClient.Stop(ctx)

		// setup api
		apiServer, err := api.NewServer(api.NewService(repo, index), &apiSecurityHandler{APIKey: config.APIKey})
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
		logger.Info("starting server", "addr", config.Addr())
		if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
			return err
		}
		logger.Info("gracefully stopped server")

		return nil
	},
}

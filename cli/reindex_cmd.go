package cli

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/spf13/cobra"
	index "github.com/ugent-library/index/es6"
	"github.com/ugent-library/projects-service/models"
	"github.com/ugent-library/projects-service/repositories"
	"github.com/ugent-library/projects-service/search/es6"
)

func init() {
	rootCmd.AddCommand(reindexCmd)
}

var reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "reindex all projects in de search index",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.TODO()
		startTime := time.Now()

		// setup repo
		repo, err := repositories.New(repositories.Config{
			Conn:   config.Repo.Conn,
			Secret: []byte(config.Repo.Secret),
		})
		if err != nil {
			return err
		}

		// setup ES client
		client, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: strings.Split(config.Search.Conn, ","),
		})

		if err != nil {
			return err
		}

		// init a switcher
		settings, err := os.ReadFile("etc/search/es6/project.json")
		if err != nil {
			return err
		}

		switcher, err := index.NewSwitcher(client, config.Search.Index, string(settings))
		if err != nil {
			return err
		}

		// get index name
		idx := switcher.Name()

		// init an indexer
		indexer, err := index.NewIndexer(client, idx, index.IndexerConfig{
			OnError: func(err error) {
				log.Printf("%+v", err)
			},
			OnIndexFailure: func(str string, err error) {
				log.Printf("%+v", err)
				log.Println("FAILED")
			},
			OnIndexSuccess: func(str string) {
				log.Println("SUCCESS")
			},
		})
		if err != nil {
			return err
		}

		// first indexing round
		repo.EachProject(ctx, func(p *models.Project) bool {
			doc := es6.NewProjectDocument(p)

			d, _ := json.Marshal(doc)
			// if err != nil {
			// 	// error
			// }

			log.Println(p.ID)

			indexer.Index(ctx, p.ID, d)
			return true
		})

		// switch the index
		switcher.Switch(ctx, config.Search.Retention)

		endTime := time.Now()

		// second indexing round
		repo.BetweenProjects(ctx, startTime, endTime, func(p *models.Project) bool {
			doc := es6.NewProjectDocument(p)

			d, _ := json.Marshal(doc)
			// if err != nil {
			// 	// error
			// }

			indexer.Index(ctx, p.ID, d)
			return true
		})

		indexer.Close(ctx)

		return nil
	},
}

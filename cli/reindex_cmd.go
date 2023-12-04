package cli

import (
	"context"
	"encoding/json"
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
			Conn: config.Repo.Conn,
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
		settings, err := os.ReadFile("etc/es6/settings.json")
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
				logger.Errorf("Indexer error: %w", err)
			},
			OnIndexFailure: func(str string, err error) {
				logger.Errorf("Failed indexing document: %s, %w", str, err)
			},
			OnIndexSuccess: func(str string) {
				logger.Infof("Indexed document: %s", str)
			},
		})
		if err != nil {
			return err
		}

		// first indexing round
		repo.EachProject(ctx, func(p *models.Project) bool {
			doc := es6.NewProjectDocument(p)

			d, _ := json.Marshal(doc)

			indexer.Index(ctx, p.ID, d)
			return true
		})

		// switch the index
		switcher.Switch(ctx, config.Search.Retention)

		endTime := time.Now()

		// second indexing round
		repo.EachProjectBetween(ctx, startTime, endTime, func(p *models.Project) bool {
			doc := es6.NewProjectDocument(p)

			d, _ := json.Marshal(doc)

			indexer.Index(ctx, p.ID, d)
			return true
		})

		indexer.Close(ctx)

		return nil
	},
}

package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reindexCmd)
}

var reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "reindex all projects in de search index",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

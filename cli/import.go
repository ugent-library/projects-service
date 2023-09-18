package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use: "import",
}

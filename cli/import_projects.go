package cli

import (
	"log"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/ugent-library/projects/gismo"
)

func init() {
	importCmd.AddCommand(importProjectsCmd)
}

var importProjectsCmd = &cobra.Command{
	Use: "projects",
	Run: func(cmd *cobra.Command, args []string) {
		c := gismo.Config{
			URL: "nats://localhost:4222",
		}
		log.Println("listening...")
		err := gismo.Listen(c)
		if err != nil {
			log.Printf("%+v", err)
		}

		runtime.Goexit()
	},
}

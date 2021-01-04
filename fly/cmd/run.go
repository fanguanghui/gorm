package cmd

import (
	"gormfly/entity/lib"
	"log"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use `dbmodel run` to generate model file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.GenerateModel(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

package cmd

import (
	"gormfly/entity/lib"
	"log"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use `gormEntity run` to generate struct file",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := lib.GetDbConf()
		if err != nil {
			log.Fatalln(err)
		}
		if err := lib.GenerateModel(conf); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

package cmd

import (
	"fmt"
	"gormfly/dao/lib"
	"log"

	"github.com/spf13/cobra"
)

var confCmd = &cobra.Command{
	Use:   "conf",
	Short: fmt.Sprintf("use `gormDao conf` to generate %s", lib.ConfigFile),
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.GenerateConf(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}

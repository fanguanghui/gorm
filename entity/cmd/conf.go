package cmd

import (
	"fmt"
	"gormfly/entity/lib"
	"log"

	"github.com/spf13/cobra"
)

var confCmd = &cobra.Command{
	Use:   "conf",
	Short: fmt.Sprintf("use `gormEntity conf` to generate %s", lib.DbConfFile),
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.GenerateDbConf(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
}

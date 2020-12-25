package cmd

import (
	"dbmodel/lib"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var confCmd = &cobra.Command{
	Use:   "conf",
	Short: fmt.Sprintf("use `dbmodel conf` to generate %s", lib.DbConfFile),
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.GenerateDbConf(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

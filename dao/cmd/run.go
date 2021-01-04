package cmd

import (
	"gormfly/dao/lib"
	"log"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Use `gormDao run` to generate Dao file",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := lib.GetConf()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(conf)

		p := lib.NewParser(conf.Input)
		gen := lib.NewGenerator(conf.Output).SetImportPkg(conf.ImportPkgs).SetLogName(conf.LogName)
		if conf.TransformErr {
			gen = gen.TransformError()
		}
		if err := gen.ParserAST(p, conf.Structs).Generate().Format().Flush(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

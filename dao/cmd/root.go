package cmd

import (
	"fmt"
	"gormfly/dao/lib"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var conf = lib.Config{}

var rootCmd = &cobra.Command{
	Use:   "gormDao",
	Short: "Use `gormDao` generate golang Dao file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.CheckConf(&conf); err != nil {
			log.Fatalln(err)
		}
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&conf.Input, "input", ",model", "[Required] The name of the input file dir")
	rootCmd.Flags().StringVar(&conf.Output, "output", ",model", "[Option] The name of the output file dir")
	rootCmd.Flags().StringArrayVar(&conf.Structs, "structs", nil, "[Required] The name of schema structs to generate structs for, comma seperated")
	rootCmd.Flags().StringArrayVar(&conf.Imports, "imports", nil, "[Required] The name of the import  to import package")
	rootCmd.Flags().StringVar(&conf.LogName, "logName", "", "[Option] The name of log db error")
	rootCmd.Flags().BoolVar(&conf.TransformErr, "transformErr", false, "[Option] The name of transform db err")
	//rootCmd.Flags().Usage()
	//rootCmd.Flags().Parsed()
}

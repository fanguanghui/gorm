package cmd

import (
	"fmt"
	"gormfly/entity/lib"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var conf = lib.DbConf{}

var rootCmd = &cobra.Command{
	Use:   "gormEntity",
	Short: "Use `gormEntity` generate golang struct file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.CheckDbConf(&conf); err != nil {
			log.Fatalln(err)
		}
		if err := lib.GenerateModel(&conf); err != nil {
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
	rootCmd.Flags().StringVar(&conf.Host, "host", "localhost", "A help for foo")
	rootCmd.Flags().IntVar(&conf.Port, "port", 3306, "A help for foo")
	rootCmd.Flags().StringVar(&conf.User, "user", "root", "A help for foo")
	rootCmd.Flags().StringVar(&conf.Password, "pwd", "", "A help for foo")
	rootCmd.Flags().StringVar(&conf.Database, "db", "", "A help for foo")
	rootCmd.Flags().StringVar(&conf.TableName, "table", "", "A help for foo")
	rootCmd.Flags().StringVar(&conf.PackageName, "pkg", "", "A help for foo")
	rootCmd.Flags().BoolVar(&conf.StructSorted, "sort", false, "A help for foo")
	rootCmd.Flags().BoolVar(&conf.GormAnnotation, "gorm", true, "A help for foo")
	rootCmd.Flags().BoolVar(&conf.JsonAnnotation, "json", true, "A help for foo")
	rootCmd.Flags().BoolVar(&conf.XmlAnnotation, "xml", false, "A help for foo")
	rootCmd.Flags().BoolVar(&conf.FakerAnnotation, "faker", false, "A help for foo")
	rootCmd.Flags().BoolVar(&conf.GureguTypes, "guregu", false, "A help for foo")
}

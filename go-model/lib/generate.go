package lib

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateModel() error {
	conf, err := GetDbConf()
	if err != nil {
		return err
	}
	db, err := NewDbConnect(conf.Db.Host, conf.Db.Port, conf.Db.User, conf.Db.Password, conf.Db.Database)
	if err != nil {
		return err
	}
	tables, err := GetTablesFromDb(db, conf.Db.Database, conf.Db.TableName)
	if err != nil {
		return err
	}
	for _, value := range tables {
		columns, err := GetColumnsFromDb(db, conf.Db.Database, value)
		if err != nil {
			return err
		}
		modelByte, err := GenerateModelByte(
			*columns,
			conf.Db.PackageName,
			value,
			conf.Db.GormAnnotation,
			conf.Db.JSONAnnotation,
			conf.Db.XMLAnnotation,
			conf.Db.FakerAnnotation,
			conf.Db.GureguTypes,
			conf.Db.StructSorted)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%s/%s.go", conf.Db.PackageName, value)
		err = GenerateFile(filePath, modelByte)
		if err != nil {
			return err
		}
	}

	return err
}

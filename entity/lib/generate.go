package lib

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GenerateModel(conf *DbConf) error {
	db, err := NewDbConnect(conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
	if err != nil {
		return err
	}
	tables, err := GetTablesFromDb(db, conf.Database, conf.TableName)
	if err != nil {
		return err
	}
	for _, value := range tables {
		columns, err := GetColumnsFromDb(db, conf.Database, value)
		if err != nil {
			return err
		}
		modelByte, err := GenerateModelByte(
			*columns,
			conf.PackageName,
			value,
			conf.GormAnnotation,
			conf.JsonAnnotation,
			conf.XmlAnnotation,
			conf.FakerAnnotation,
			conf.GureguTypes,
			conf.StructSorted)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%s/%s.go", conf.PackageName, value)
		err = GenerateFile(filePath, modelByte)
		if err != nil {
			return err
		}
	}

	return err
}

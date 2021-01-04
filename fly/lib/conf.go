package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var DbConfFile = "./conf.json"

var DbConfInfo = `{
  "db": {
    "host": "localhost",
    "port": 3306,
    "user": "root",
    "password": "password",
    "database": "database",
    "tableName": "tableName",
    "packageName": "packageName",
    "structSorted": false,
    "gormAnnotation": true,
    "jsonAnnotation": true,
    "xmlAnnotation": false,
    "fakerAnnotation": false,
    "gureguTypes": false
  }
}`

type DbConf struct {
	Db struct {
		Host            string `json:"host"`
		Port            int    `json:"port"`
		User            string `json:"user"`
		Password        string `json:"password"`
		Database        string `json:"database"`
		TableName       string `json:"tableName"`
		PackageName     string `json:"packageName"`
		StructSorted    bool   `json:"struct_sorted"`
		GormAnnotation  bool   `json:"gormAnnotation"`
		JSONAnnotation  bool   `json:"jsonAnnotation"`
		XMLAnnotation   bool   `json:"xmlAnnotation"`
		FakerAnnotation bool   `json:"fakerAnnotation"`
		GureguTypes     bool   `json:"gureguTypes"`
	} `json:"db"`
}

func GetDbConf() (*DbConf, error) {
	confFile, err := ioutil.ReadFile(DbConfFile)
	if err != nil {
		return nil, err
	}
	dbConf := DbConf{}
	if err = json.Unmarshal(confFile, &dbConf); err != nil {
		return nil, err
	}
	if dbConf.Db.Host == "" || dbConf.Db.Database == "" {
		err = errors.New("The host database cant be empty\n")
		return nil, err
	}
	if dbConf.Db.Port == 0 {
		dbConf.Db.Port = 3306
	}
	if dbConf.Db.PackageName == "" {
		dbConf.Db.PackageName = "model"
	}
	return &dbConf, nil
}

func GenerateDbConf() error {
	return GenerateFile(DbConfFile, []byte(DbConfInfo))
}

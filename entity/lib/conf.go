package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var DbConfFile = "./conf.json"

var DbConfInfo = `{
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
}`

type DbConf struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	User            string `json:"user"`
	Password        string `json:"password"`
	Database        string `json:"database"`
	TableName       string `json:"tableName"`
	PackageName     string `json:"packageName"`
	StructSorted    bool   `json:"struct_sorted"`
	GormAnnotation  bool   `json:"gormAnnotation"`
	JsonAnnotation  bool   `json:"jsonAnnotation"`
	XmlAnnotation   bool   `json:"xmlAnnotation"`
	FakerAnnotation bool   `json:"fakerAnnotation"`
	GureguTypes     bool   `json:"gureguTypes"`
}

func GetDbConf() (*DbConf, error) {
	confFile, err := ioutil.ReadFile(DbConfFile)
	if err != nil {
		return nil, err
	}
	dbConf := &DbConf{}
	if err = json.Unmarshal(confFile, &dbConf); err != nil {
		return nil, err
	}
	if err = CheckDbConf(dbConf); err != nil {
		return nil, err
	}
	return dbConf, err
}

func CheckDbConf(dbConf *DbConf) error {
	if dbConf.Host == "" || dbConf.Database == "" {
		return errors.New("The host database cant be empty\n")
	}
	if dbConf.Port == 0 {
		dbConf.Port = 3306
	}
	if len(dbConf.PackageName) == 0 {
		dbConf.PackageName = "model"
	}
	return nil
}

func GenerateDbConf() error {
	return GenerateFile(DbConfFile, []byte(DbConfInfo))
}

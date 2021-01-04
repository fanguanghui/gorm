package lib

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var ConfigFile = "./conf.yaml"

type Config struct {
	Input        string      `yaml:"input"`
	Output       string      `yaml:"output"`
	Imports      []string    `yaml:"imports"`
	Structs      []string    `yaml:"structs"`
	LogName      string      `yaml:"logName"`
	TransformErr bool        `yaml:"transformErr"`
	ImportPkgs   []ImportPkg `yaml:"-"`
}

func GetConf() (*Config, error) {
	confFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return nil, err
	}
	dbConf := &Config{}
	if err = yaml.Unmarshal(confFile, &dbConf); err != nil {
		return nil, err
	}
	if err = CheckConf(dbConf); err != nil {
		return nil, err
	}
	return dbConf, err
}

func CheckConf(conf *Config) error {
	if len(conf.Input) == 0 {
		return errors.New("The Input cant be empty\n")
	}
	if len(conf.Output) == 0 {
		conf.Output = conf.Input
	}

	for _, v := range conf.Imports {
		conf.ImportPkgs = append(conf.ImportPkgs, ImportPkg{
			Pkg: v,
		})
	}
	return nil
}

func GenerateConf() error {
	conf := Config{
		Input:        "./model",
		Output:       "./model",
		Imports:      []string{"github.com/jinzhu/gorm"},
		Structs:      nil,
		LogName:      "logName",
		TransformErr: false,
	}
	bytes, err := yaml.Marshal(&conf)
	if err != nil {
		return err
	}
	return GenerateFile(ConfigFile, bytes)
}

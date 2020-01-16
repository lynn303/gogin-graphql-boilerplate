/**
 * @Time  : 2020-01-16 10:08
 * @Author: Lynn
 * @File  : config
 * @Description:
 * @History:
 *  1.[2020-01-16-10:08] new created
 */
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type mysqlConfig struct {
	Database string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type config struct {
	MysqlConfig mysqlConfig `json:"mysql"`
}

func InitConfig() (e *config, err error) {
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		log.Printf("yamlFile.Read err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &e)
	if err != nil {
		log.Fatal("Unmarshal:%v", err)
	}
	return e, err
}

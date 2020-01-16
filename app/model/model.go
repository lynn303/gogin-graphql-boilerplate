/**
 * @Time  : 2020-01-16 18:24
 * @Author: Lynn
 * @File  : model
 * @Description:
 * @History:
 *  1.[2020-01-16-18:24] new created
 */
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gogin-graphql-boilerplate/config"
)

var db *gorm.DB

func init() {
	config, _ := config.InitConfig()
	var err error
	db, err = gorm.Open(config.MysqlConfig.Name, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MysqlConfig.User,
		config.MysqlConfig.Password,
		config.MysqlConfig.Host,
		config.MysqlConfig.Port,
		config.MysqlConfig.Database))
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(2)
	//data migration
	migrations(db)
	//init data
	initData()
}

func DB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}

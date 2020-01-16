/**
 * @Time  : 2020-01-16 10:08
 * @Author: Lynn
 * @File  : migration
 * @Description:
 * @History:
 *  1.[2020-01-16-10:08] new created
 */
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func migrations(db *gorm.DB) {
	tableList := []interface{}{
		&User{}, &Contact{},
		//TODO 添加其他的数据库model.....
	}
	i := 0
	for _, u := range tableList {
		if !db.HasTable(u) {
			db.AutoMigrate(u)
			i++
		}
	}
	fmt.Println("[data migration] success, migration table num:", i)
}

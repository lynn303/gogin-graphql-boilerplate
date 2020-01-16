/**
 * @Time  : 2020-01-16 10:02
 * @Author: Lynn
 * @File  : contact
 * @Description:
 * @History:
 *  1.[2020-01-16-10:02] new created
 */
package model

import "github.com/jinzhu/gorm"

type Contact struct {
	gorm.Model
	UserId   uint   `json:"userId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Position string `json:"position"`
	QQ       string `json:"qq"`
	Wechat   string `json:"wechat"`
}

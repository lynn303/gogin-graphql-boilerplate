/**
 * @Time  : 2020-01-16 10:00
 * @Author: Lynn
 * @File  : user
 * @Description:
 * @History:
 *  1.[2020-01-16-10:00] new created
 */
package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	HeadUrl  string `json:"headUrl"`
	Active   bool   `json:"active"`
}

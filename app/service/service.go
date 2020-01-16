/**
 * @Time  : 2020-01-16 10:06
 * @Author: Lynn
 * @File  : service
 * @Description:
 * @History:
 *  1.[2020-01-16-10:06] new created
 */
package service

import "time"

type UserInput struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	HeadUrl  string `json:"headUrl"`
	Active   bool   `json:"active"`
}
type UserRst struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"userName"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	HeadUrl   string    `json:"headUrl"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
type ContactInput struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"userId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Position string `json:"position"`
	QQ       string `json:"qq"`
	Wechat   string `json:"wechat"`
}
type ContactRst struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"userId"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Position  string    `json:"position"`
	QQ        string    `json:"qq"`
	Wechat    string    `json:"wechat"`
	CreatedAt time.Time `json:"createdAt"`
}

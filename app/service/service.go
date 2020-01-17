/**
 * @Time  : 2020-01-16 10:06
 * @Author: Lynn
 * @File  : service
 * @Description:
 * @History:
 *  1.[2020-01-16-10:06] new created
 */
package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"gogin-graphql-boilerplate/app/app_utils"
	"gogin-graphql-boilerplate/app/model"
	"gogin-graphql-boilerplate/utils"
	"time"
)

type userInput struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	HeadUrl  string `json:"headUrl"`
	Active   bool   `json:"active"`
}
type userRst struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"userName"`
	HeadUrl   string    `json:"headUrl"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
type contactInput struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"userId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Position string `json:"position"`
	QQ       string `json:"qq"`
	Wechat   string `json:"wechat"`
}
type contactRst struct {
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

func AopUser(data map[string]interface{}) (app_utils.Result, error) {
	var (
		user  model.User
		input userInput
	)
	if err := mapstructure.Decode(data, &input); err != nil {
		return app_utils.HsResult(false, errors.New("params error"))
	}
	_salt := utils.GenerateSmsCode(32)
	_password := utils.MD5(input.Password, _salt)
	if input.ID > 0 {
		//update
		if err := model.DB().Where("id = ?", input.ID).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return app_utils.HsResult(false, errors.New("user is not exists"))
			}
			return app_utils.HsResult(false, errors.New("failed"))
		}
		user.UserName = input.UserName
		user.Password = _password
		user.Salt = _salt
		user.HeadUrl = input.HeadUrl
		if err := model.DB().Save(&user).Error; err != nil {
			return app_utils.HsResult(false, errors.New("update user error"))
		}
	} else {
		//new
		userInput := model.User{
			UserName: input.UserName,
			Password: _password,
			Salt:     _salt,
			HeadUrl:  input.HeadUrl,
			Active:   false,
		}
		if err := model.DB().Create(&userInput).Error; err != nil {
			return app_utils.HsResult(false, errors.New("new user error"))
		}
	}
	return app_utils.HsResult(true, nil)
}

func DelUser(id uint) (app_utils.Result, error) {
	tx := model.DB().Begin()
	if err := tx.Where("user_id = ?", id).Delete(&model.Contact{}).Error; err != nil {
		tx.Callback()
		return app_utils.HsResult(false, errors.New("delete contact error"))
	}
	if err := tx.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		tx.Callback()
		return app_utils.HsResult(false, errors.New("delete user error"))
	}
	return app_utils.HsResult(true, nil)
}

func GetUserList() ([]userRst, error) {
	var (
		list     []model.User
		outputes []userRst
	)
	if err := model.DB().Order("created_at desc").Find(&list).Error; err != nil {
		return outputes, nil
	}
	for _, item := range list {
		output := userRst{
			ID:        item.ID,
			UserName:  item.UserName,
			HeadUrl:   item.HeadUrl,
			Active:    item.Active,
			CreatedAt: item.CreatedAt,
		}
		outputes = append(outputes, output)
	}
	return outputes, nil
}

/**
 * @Time  : 2020-01-16 18:26
 * @Author: Lynn
 * @File  : initdata
 * @Description:
 * @History:
 *  1.[2020-01-16-18:26] new created
 */
package model

import (
	"fmt"
	"gogin-graphql-boilerplate/utils"
)

func initData() {
	initUser()
}

func initUser() {
	var (
		count int
	)
	if err := DB().Model(&User{}).Count(&count).Error; err != nil {
		fmt.Println("init user data error:", err)
	}
	if count > 0 {
		return
	}
	defaultPassword := "123456"
	salt := utils.GenerateSmsCode(32)
	password := utils.MD5(defaultPassword, salt)
	user := User{
		UserName: "admin",
		Password: password,
		Salt:     salt,
		HeadUrl:  "https://wx.qlogo.cn/mmopen/vi_32/MxIKVrOfvGo6OUloRianUcWreEUqOd0BncnkDQLHAHolCxvxWgic4Zq8lUgYLoof9xkAvRaOAgwH43KOBYbthFiag/132",
		Active:   false,
	}
	if err := DB().Create(&user).Error; err != nil {
		fmt.Println("init user data error:", err)
	}
}

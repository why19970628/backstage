//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:11:17
 * @LastEditors: congz
 * @LastEditTime: 2020-09-21 14:10:10
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	Nickname       string `gorm:"unique"`
	Status         string
	Limit          int
	Avatar         string `gorm:"size:1000"`
}

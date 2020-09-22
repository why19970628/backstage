//Package model ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:11:17
 * @LastEditors: congz
 * @LastEditTime: 2020-09-22 13:31:16
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
	Limit          uint32
	Avatar         string `gorm:"size:1000"`
}

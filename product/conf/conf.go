//Package conf ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-09-20 11:07:14
 */
package conf

import (
	"os"
	"product/model"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

}

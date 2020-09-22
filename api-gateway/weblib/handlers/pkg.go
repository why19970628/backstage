// Package handlers ..
/*
 * @Descripttion:处理函数
 * @Author: congz
 * @Date: 2020-09-20 11:33:37
 * @LastEditors: congz
 * @LastEditTime: 2020-09-22 14:01:28
 */
package handlers

//PanicIfError 包装错误函数
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

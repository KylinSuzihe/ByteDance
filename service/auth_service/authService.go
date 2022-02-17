package auth_service

import (
	"demo1/models"
	"demo1/pkg/e"
	"demo1/pkg/util"
)

// Login 用户登录
// 作者：邱港，刘丰利
func Login(username, password string) (string, e.ErrNo) {
	var userId string
	var errno e.ErrNo

	// 这里不管是不是参数错误，返回的就是密码错误
	if !util.CheckUserName(username) || !util.CheckPassword(password) {
		userId = ""
		errno = e.WrongPassword
	} else {
		userId, errno = models.CheckAuth(username, password)
	}
	return userId, errno
}

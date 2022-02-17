package models

import (
	"demo1/pkg/e"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func (Member) TableName() string {
	return "member"
}

// CheckAuth 用户登录
// 作者：邱港，刘丰利
func CheckAuth(username, password string) (userId string, errno e.ErrNo) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in PanicError", r)
			userId = ""
			errno = e.WrongPassword
		}
	}()

	var member Member
	err := db.Select("user_id").
		Where("username = ? AND password = ? AND deleted = ?", username, password, 0).
		First(&member).Error

	// 如果发生了错误
	if err != nil {
		// 如果发生的错误不是用户未找到，返回错误
		if err != gorm.ErrRecordNotFound {
			return "", e.UnknownError
		}
		// 如果发生的错误是用户未找到，返回0，nil
		return "", e.WrongPassword
	}
	// 没有发生错误，返回用户id
	return strconv.FormatInt(member.UserId, 10), e.OK
}

package controller

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
	"demo1/service/auth_service"
	"demo1/service/member_service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Login 用户登录
// 作者：邱港，刘丰利
func Login(c *gin.Context) {
	var request types.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, types.LoginResponse{
			Code: e.ParamInvalid,
			Data: struct{ UserID string }{UserID: ""},
		})
		return
	}
	userId, errno := auth_service.Login(request.Username, request.Password)

	if errno == e.ParamInvalid {
		c.JSON(http.StatusBadRequest, types.LoginResponse{
			Code: errno,
			Data: struct{ UserID string }{UserID: userId},
		})
		return
	}
	session := sessions.Default(c)
	session.Set("userId", userId)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusOK, types.LoginResponse{
			Code: e.UnknownError,
			Data: struct{ UserID string }{UserID: ""},
		})
		return
	}
	c.JSON(http.StatusOK, types.LoginResponse{
		Code: errno,
		Data: struct{ UserID string }{UserID: userId},
	})
}

// Logout 用户登出
// 作者：邱港，刘丰利
func Logout(c *gin.Context) {
	value, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(http.StatusOK, types.LogoutResponse{
			Code: e.LoginRequired,
		})
		return
	}
	session := sessions.Default(c)
	session.Clear()

	if err = session.Save(); err != nil {
		c.JSON(http.StatusOK, types.LogoutResponse{
			Code: e.UnknownError,
		})
		return
	}
	cookie := http.Cookie{
		Name:     "camp-session",
		Value:    value,
		Path:     "/",
		Domain:   "",
		Expires:  time.Time{},
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: false,
	}
	http.SetCookie(c.Writer, &cookie)
	c.JSON(http.StatusOK, types.LogoutResponse{
		Code: e.OK,
	})
}

// WhoAmI 测试登录是否成功
// 作者：邱港，刘丰利
func WhoAmI(c *gin.Context) {
	value, err := c.Cookie("camp-session")
	if err != nil {
		c.JSON(http.StatusOK, types.LogoutResponse{
			Code: e.LoginRequired,
		})
		return
	}
	session := sessions.Default(c)
	userId := session.Get("userId").(string)
	member, errno := member_service.GetMember(userId)

	if member == nil {
		c.JSON(http.StatusOK, types.WhoAmIResponse{
			Code: errno,
			Data: types.TMember{},
		})

		session.Clear()
		if err = session.Save(); err != nil {
			c.JSON(http.StatusOK, types.LogoutResponse{
				Code: e.UnknownError,
			})
			return
		}
		cookie := http.Cookie{
			Name:     "camp-session",
			Value:    value,
			Path:     "/",
			Domain:   "",
			Expires:  time.Time{},
			MaxAge:   -1,
			Secure:   false,
			HttpOnly: false,
		}
		http.SetCookie(c.Writer, &cookie)
		return
	}

	var whoami = types.WhoAmIResponse{
		Code: errno,
		Data: types.TMember{
			UserID:   strconv.FormatInt(member.UserId, 10),
			Nickname: member.Nickname,
			Username: member.Username,
			UserType: member.UserType,
		},
	}
	c.JSON(http.StatusOK, whoami)
}

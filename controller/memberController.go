package controller

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
	"demo1/service/member_service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateMember(context *gin.Context) {
	var createMemberRequest types.CreateMemberRequest
	if err := context.ShouldBindJSON(&createMemberRequest); err != nil {
		context.JSON(http.StatusBadRequest, types.CreateMemberResponse{
			Code: e.ParamInvalid,
			Data: struct{ UserID string }{UserID: ""},
		})
		return
	}
	// 创建用户会携带cookie，需要检查权限
	value, err := context.Cookie("camp-session")
	if err != nil {
		context.JSON(http.StatusOK, types.CreateMemberResponse{
			Code: e.LoginRequired,
			Data: struct{ UserID string }{UserID: ""},
		})
		return
	}
	session := sessions.Default(context)
	userId := session.Get("userId").(string)
	member, errno := member_service.GetMember(userId)

	// 如果根据userid查询出来的member不存在，返回需要登录并将cookie中的内容清除
	if member == nil {
		context.JSON(http.StatusOK, types.CreateMemberResponse{
			Code: e.LoginRequired,
			Data: struct{ UserID string }{UserID: ""},
		})
		session.Clear()
		if err = session.Save(); err != nil {
			context.JSON(http.StatusOK, types.CreateMemberResponse{
				Code: e.UnknownError,
				Data: struct{ UserID string }{UserID: ""},
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
		http.SetCookie(context.Writer, &cookie)
		return
	}

	if member.Username != "JudgeAdmin" || member.Password != "JudgePassword2022" {
		context.JSON(http.StatusOK, types.CreateMemberResponse{
			Code: e.PermDenied,
			Data: struct{ UserID string }{UserID: ""},
		})
		return
	}

	userId, errno = member_service.CreateMember(&createMemberRequest)

	context.JSON(http.StatusOK, types.CreateMemberResponse{
		Code: errno,
		Data: struct{ UserID string }{UserID: userId},
	})
	return
}

func GetMember(context *gin.Context) {
	var getMemberRequest types.GetMemberRequest
	var exist bool
	getMemberRequest.UserID, exist = context.GetQuery("UserID")

	// 如果不存在参数
	if !exist {
		context.JSON(http.StatusBadRequest, types.GetMemberResponse{
			Code: e.ParamInvalid,
			Data: types.TMember{},
		})
		return
	}

	member, errno := member_service.GetMember(getMemberRequest.UserID)

	// 如果用户不存在
	if member == nil {
		context.JSON(http.StatusOK, types.GetMemberResponse{
			Code: errno,
			Data: types.TMember{},
		})
		return
	}

	context.JSON(http.StatusOK, types.GetMemberResponse{
		Code: errno,
		Data: types.TMember{
			UserID:   strconv.FormatInt(member.UserId, 10),
			Nickname: member.Nickname,
			Username: member.Username,
			UserType: member.UserType,
		},
	})
	return
}

func GetMemberList(context *gin.Context) {
	offset, exist1 := context.GetQuery("Offset")
	limit, exist2 := context.GetQuery("Limit")
	if !exist1 || !exist2 {
		context.JSON(http.StatusBadRequest, types.GetMemberListResponse{
			Code: e.ParamInvalid,
			Data: struct{ MemberList []types.TMember }{MemberList: nil},
		})
		return
	}

	var getMemberListRequest types.GetMemberListRequest
	var err1, err2 error
	getMemberListRequest.Offset, err1 = strconv.Atoi(offset)
	getMemberListRequest.Limit, err2 = strconv.Atoi(limit)

	if err1 != nil || err2 != nil {
		context.JSON(http.StatusBadRequest, types.GetMemberListResponse{
			Code: e.ParamInvalid,
			Data: struct{ MemberList []types.TMember }{MemberList: nil},
		})
		return
	}

	memberList, errno := member_service.GetMemberList(getMemberListRequest.Offset, getMemberListRequest.Limit)

	context.JSON(http.StatusOK, types.GetMemberListResponse{
		Code: errno,
		Data: struct{ MemberList []types.TMember }{MemberList: memberList},
	})
	return
}

func UpdateMember(context *gin.Context) {
	var updateMemberRequest types.UpdateMemberRequest
	if err := context.ShouldBindJSON(&updateMemberRequest); err != nil {
		context.JSON(http.StatusBadRequest, types.UpdateMemberResponse{
			Code: e.ParamInvalid,
		})
		return
	}

	errno := member_service.UpdateMember(updateMemberRequest.UserID, updateMemberRequest.Nickname)

	context.JSON(http.StatusOK, types.UpdateMemberResponse{
		Code: errno,
	})
	return
}

func DeleteMember(context *gin.Context) {
	var deletedMemberRequest types.DeleteMemberRequest
	if err := context.ShouldBindJSON(&deletedMemberRequest); err != nil {
		context.JSON(http.StatusBadRequest, types.DeleteMemberResponse{
			Code: e.ParamInvalid,
		})
		return
	}

	errno := member_service.DeletedMember(deletedMemberRequest.UserID)
	context.JSON(http.StatusOK, types.DeleteMemberResponse{
		Code: errno,
	})
	return
}

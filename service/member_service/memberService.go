package member_service

import (
	"demo1/models"
	"demo1/pkg/e"
	"demo1/pkg/types"
	"demo1/pkg/util"
	"strconv"
)

func CreateMember(createMemberRequest *types.CreateMemberRequest) (string, e.ErrNo) {
	nickname := createMemberRequest.Nickname
	username := createMemberRequest.Username
	password := createMemberRequest.Password
	userType := createMemberRequest.UserType

	// 验证userType是否合法
	if !util.CheckUserType(userType) {
		return "", e.ParamInvalid
	}

	// 验证 Nickname，username和password
	if !util.CheckNickName(nickname) || !util.CheckUserName(username) || !util.CheckPassword(password) {
		return "", e.ParamInvalid
	}

	//验证username是否重复
	if models.ExistMemberByUsername(username) {
		return "", e.UserHasExisted
	}

	var member = &models.Member{
		UserId:   0,
		Nickname: nickname,
		Username: username,
		Password: password,
		UserType: userType,
	}
	//调用dao创建用户
	errno := models.CreateMember(member)

	return strconv.FormatInt(member.UserId, 10), errno
}

func GetMember(userId string) (member *models.Member, errno e.ErrNo) {
	userId1, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, e.ParamInvalid
	}
	return models.GetMember(userId1)
}

func GetMemberList(offset int, limit int) ([]types.TMember, e.ErrNo) {
	if offset < 0 || limit < 0 {
		return nil, e.ParamInvalid
	}
	memberList, errno := models.GetMemberList(offset, limit)
	if errno != e.OK || memberList == nil || len(memberList) == 0 {
		return nil, errno
	}
	return memberList, errno
}

func UpdateMember(userId string, nickname string) e.ErrNo {
	if !util.CheckNickName(nickname) {
		return e.ParamInvalid
	}
	userId1, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return e.ParamInvalid
	}
	return models.UpdateMember(userId1, nickname)
}

func DeletedMember(userId string) e.ErrNo {
	userId1, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return e.ParamInvalid
	}
	return models.DeleteMember(userId1)
}

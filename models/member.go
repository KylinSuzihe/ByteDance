package models

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
	"fmt"
	"gorm.io/gorm"
)

func CreateMember(member *Member) (errno e.ErrNo) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in PanicError", r)
			errno = e.UnknownError
		}
	}()

	err := db.Create(member).Error
	if err != nil {
		errno = e.UnknownError
		return
	}

	errno = e.OK
	return
}

func GetMember(userId int64) (*Member, e.ErrNo) {
	var member Member
	err := db.Where("user_id = ?", userId).
		First(&member).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, e.UserNotExisted
		}
		fmt.Printf("member.GetMember err %v", err)
		return nil, e.UnknownError
	}

	if member.Deleted == 1 {
		return nil, e.UserHasDeleted
	}
	return &member, e.OK
}

func GetMemberList(offset int, limit int) ([]types.TMember, e.ErrNo) {
	var memberList []types.TMember
	err := db.Where("deleted = ?", 0).
		Limit(limit).
		Offset(offset).
		Find(&memberList).Error

	if err != nil {
		return nil, e.UnknownError
	}

	return memberList, e.OK
}

func UpdateMember(userId int64, nickname string) e.ErrNo {
	var member Member
	err := db.Select("nickname").
		Where("user_id = ? AND deleted = ?", userId, 0).
		Find(&member).Error

	if err != nil {
		return e.UserNotExisted
	}

	if member.Nickname == nickname {
		return e.OK
	}

	tx := db.Model(&Member{}).
		Where("user_id = ? AND deleted = ?", userId, 0).
		Update("nickname", nickname)
	if tx.Error != nil {
		return e.UnknownError
	}
	if tx.RowsAffected == 0 {
		return e.UserNotExisted
	}
	return e.OK
}

func DeleteMember(userId int64) e.ErrNo {
	tx := db.Model(&Member{}).
		Where("user_id = ? AND deleted = ?", userId, 0).
		Update("deleted", 1)
	if tx.Error != nil {
		return e.UnknownError
	}
	if tx.RowsAffected == 0 {
		return e.UserNotExisted
	}
	return e.OK
}

func ExistMemberByUsername(username string) bool {
	var member Member
	err := db.Select("username").
		Where("username = ? AND deleted = ?", username, 0).
		First(&member).Error

	if err != nil {
		return false
	}
	return member.Username != ""
}

func ExistMemberByUserId(userId int64, userType types.UserType) bool {
	var member Member
	err := db.Select("user_id").
		Where("user_id = ? AND user_type = ? AND deleted = ?", userId, userType, 0).
		First(&member).Error

	if err != nil {
		return false
	}
	return member.UserId != 0
}

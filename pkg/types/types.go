package types

import (
	"demo1/pkg/e"
)

type TMember struct {
	UserID   string
	Nickname string
	Username string
	UserType UserType
}

type TCourse struct {
	CourseID  string
	Name      string
	TeacherID string
}

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)

func (TCourse) TableName() string {
	return "course"
}

func (TMember) TableName() string {
	return "member"
}

// 系统内置管理员账号
// 账号名：JudgeAdmin 密码：JudgePassword2022

// 创建成员
// 参数不合法返回 ParamInvalid

// CreateMemberRequest 只有管理员才能添加
type CreateMemberRequest struct {
	Nickname string   // required，不小于 4 位 不超过 20 位
	Username string   // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string   // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
	UserType UserType // required, 枚举值
}

type CreateMemberResponse struct {
	Code e.ErrNo
	Data struct {
		UserID string // int64 范围
	}
}

// 获取成员信息

type GetMemberRequest struct {
	UserID string
}

// 如果用户已删除请返回已删除状态码，不存在请返回不存在状态码

type GetMemberResponse struct {
	Code e.ErrNo
	Data TMember
}

// 批量获取成员信息

type GetMemberListRequest struct {
	Offset int
	Limit  int
}

type GetMemberListResponse struct {
	Code e.ErrNo
	Data struct {
		MemberList []TMember
	}
}

// 更新成员信息

type UpdateMemberRequest struct {
	UserID   string
	Nickname string
}

type UpdateMemberResponse struct {
	Code e.ErrNo
}

// 删除成员信息
// 成员删除后，该成员不能够被登录且不应该不可见，ID 不可复用

type DeleteMemberRequest struct {
	UserID string
}

type DeleteMemberResponse struct {
	Code e.ErrNo
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Code e.ErrNo
	Data struct {
		UserID string
	}
}

type LogoutResponse struct {
	Code e.ErrNo
}

type WhoAmIResponse struct {
	Code e.ErrNo
	Data TMember
}

// -------------------------------------
// 排课

// CreateCourseRequest 创建课程
// Method: Post
type CreateCourseRequest struct {
	Name string
	Cap  int
}

type CreateCourseResponse struct {
	Code e.ErrNo
	Data struct {
		CourseID string
	}
}

// GetCourseRequest 获取课程
// Method: Get
type GetCourseRequest struct {
	CourseID string
}

type GetCourseResponse struct {
	Code e.ErrNo
	Data TCourse
}

// BindCourseRequest 老师绑定课程
// Method： Post
// 注：这里的 teacherID 不需要做已落库校验
// 一个老师可以绑定多个课程 , 不过，一个课程只能绑定在一个老师下面
type BindCourseRequest struct {
	CourseID  string
	TeacherID string
}

type BindCourseResponse struct {
	Code e.ErrNo
}

// UnbindCourseRequest 老师解绑课程
// Method： Post
type UnbindCourseRequest struct {
	CourseID  string
	TeacherID string
}

type UnbindCourseResponse struct {
	Code e.ErrNo
}

// GetTeacherCourseRequest 获取老师下所有课程
// Method：Get
type GetTeacherCourseRequest struct {
	TeacherID string
}

type GetTeacherCourseResponse struct {
	Code e.ErrNo
	Data struct {
		CourseList []*TCourse
	}
}

// ScheduleCourseRequest 排课求解器，使老师绑定课程的最优解， 老师有且只能绑定一个课程
// Method： Post
type ScheduleCourseRequest struct {
	TeacherCourseRelationShip map[string][]string // key 为 teacherID , val 为老师期望绑定的课程 courseID 数组
}

type ScheduleCourseResponse struct {
	Code e.ErrNo
	Data map[string]string // key 为 teacherID , val 为老师最终绑定的课程 courseID
}

type BookCourseRequest struct {
	StudentID string
	CourseID  string
}

// BookCourseResponse 课程已满返回 CourseNotAvailable
type BookCourseResponse struct {
	Code e.ErrNo
}

type GetStudentCourseRequest struct {
	StudentID string
}

type GetStudentCourseResponse struct {
	Code e.ErrNo
	Data struct {
		CourseList []TCourse
	}
}

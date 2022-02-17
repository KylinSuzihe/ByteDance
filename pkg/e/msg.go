package e

var msgFlags = map[ErrNo]string{
	OK:                 "ok",
	ParamInvalid:       "参数不合法",
	UserHasExisted:     "该 Username 已存在",
	UserHasDeleted:     "用户已删除",
	UserNotExisted:     "用户不存在",
	WrongPassword:      "密码错误",
	LoginRequired:      "用户未登录",
	CourseNotAvailable: "课程已满",
	CourseHasBound:     "课程已绑定过",
	CourseNotBind:      "课程未绑定过",
	PermDenied:         "没有操作权限",
	StudentNotExisted:  "学生不存在",
	CourseNotExisted:   "课程不存在",
	StudentHasNoCourse: "学生没有课程",
	StudentHasCourse:   "学生有课程",
	UnknownError:       "未知错误",
}

// GetMsg get e information based on Code
func GetMsg(code ErrNo) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[UnknownError]
}

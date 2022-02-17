package teacher_service

import (
	"demo1/models"
	"demo1/pkg/e"
	"demo1/pkg/types"
)

// 作者：田思润

func BindCourse(courseId string, teacherId string) e.ErrNo {
	if courseId == "" || teacherId == "" {
		return e.ParamInvalid
	}
	return models.BindCourse(courseId, teacherId)
}

func UnbindCourse(courseId string, teacherId string) e.ErrNo {
	if courseId == "" || teacherId == "" {
		return e.ParamInvalid
	}
	return models.UnbindCourse(courseId, teacherId)
}

func GetTeacherCourse(teacherId string) ([]*types.TCourse, e.ErrNo) {
	if teacherId == "" {
		return nil, e.PermDenied
	}
	return models.GetTeacherCourse(teacherId)
}

package models

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
)

// 作者：田思润

func checkCourse(courseId string) (*Course, e.ErrNo) {
	var course Course
	err := db.Debug().Where("course_id = ? AND deleted = ?", courseId, 0).Find(&course).Error

	if err != nil {
		return nil, e.UnknownError
	}
	if course.CourseId == 0 {
		return nil, e.CourseNotExisted
	}
	return &course, e.OK
}

func BindCourse(courseId, teacherId string) e.ErrNo {
	course, err := checkCourse(courseId)
	if err != 0 {
		return err
	}
	if course.TeacherId != "" {
		return e.CourseHasBound
	}

	tx := db.Model(&course).
		Where("course_id = ? AND teacher_id = ? AND deleted = ?", courseId, "", 0).
		Update("teacher_id", teacherId)
	if tx.Error != nil {
		return e.UnknownError
	}
	if tx.RowsAffected == 0 {
		return e.CourseNotBind
	}
	return e.OK
}

func UnbindCourse(courseId, teacherId string) e.ErrNo {
	course, err := checkCourse(courseId)
	if err != 0 {
		return err
	}
	if course.TeacherId == "" {
		return e.CourseNotBind
	}
	
	if course.TeacherId != teacherId {
		return e.ParamInvalid
	}
	tx := db.Model(&course).
		Where("course_id = ? AND teacher_id = ? AND deleted = ?", courseId, teacherId, 0).
		Update("teacher_id", "")
	if tx.Error != nil {
		return e.UnknownError
	}

	if tx.RowsAffected == 0 {
		return e.CourseNotBind
	}

	return e.OK
}

func GetTeacherCourse(teacherId string) ([]*types.TCourse, e.ErrNo) {
	var courseList []*types.TCourse
	err := db.Where("teacher_id = ? AND deleted = ?", teacherId, 0).Find(&courseList).Error
	if err != nil {
		return nil, e.UnknownError
	}
	return courseList, e.OK
}

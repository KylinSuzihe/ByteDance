package models

import (
	"context"
	"demo1/pkg/e"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

// 作者：田思润

func (Course) TableName() string {
	return "course"
}

func CreateCourseWithRedis(course *Course) (errno e.ErrNo) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in PanicError", r)
			tx.Rollback()
			errno = e.UnknownError
		}
	}()

	if err := tx.Error; err != nil {
		return e.UnknownError
	}

	if err := tx.Create(course).Error; err != nil {
		tx.Rollback()
		errno = e.UnknownError
		return
	}
	if err := rdb.Set(context.Background(), "course:"+strconv.FormatInt(course.CourseId, 10), course.Capacity, 0).Err(); err != nil {
		tx.Rollback()
		errno = e.UnknownError
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		errno = e.UnknownError
		fmt.Printf("%v\n", err)
		return
	}
	errno = e.OK
	return
}

func GetCourse(courseId string) (*Course, e.ErrNo) {
	var course Course
	err := db.Where("course_id = ?", courseId).First(&course).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, e.CourseNotExisted
		}
		fmt.Printf("course.GetCourse err %v", err)
		return nil, e.UnknownError
	}
	if course.Deleted == 1 {
		return nil, e.CourseNotExisted
	}
	return &course, e.OK
}

package models

import (
	"context"
	"demo1/pkg/e"
	"demo1/pkg/types"
	"fmt"
	"gorm.io/gorm"
)

func (CourseBook) TableName() string {
	return "course_book"
}

// BookCourse 利用update语句的锁机制/**
func BookCourse(studentId, courseId string) (errno e.ErrNo) {
	// 假设studentId和courseId都存在
	var course Course
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			println("models.BookCourse 发生错误")
			errno = e.UnknownError
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return e.UnknownError
	}

	tx = tx.Model(&course).
		Where("course_id = ? AND capacity > 0 AND deleted = 0", courseId).
		UpdateColumn("capacity", gorm.Expr("capacity - ?", 1))

	if err := tx.Error; err != nil {
		fmt.Printf("err:%v", err)
		tx.Rollback()
		return e.UnknownError
	}

	if rows := tx.RowsAffected; rows == 0 {
		tx.Commit()
		return e.CourseNotAvailable
	}

	if err := db.Model(&CourseBook{}).Create(map[string]interface{}{
		"student_id": studentId, "course_id": courseId, "deleted": 0,
	}).Error; err != nil {
		tx.Rollback()
		return e.UnknownError
	}

	if err := tx.Commit().Error; err != nil {
		return e.UnknownError
	}
	errno = e.OK
	return errno
}

func HaveRemainderCapacity(courseCapacityKey string) (ret bool) {
	remain, err := rdb.Do(context.Background(), "get", courseCapacityKey).Int64()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("出现错误")
			ret = false
		}
	}()
	if err != nil {
		return false
	}
	ret = remain > 0
	return ret
}

func DecrCapacity(courseCapacityKey string) (ret int64) {
	var err error
	ret, err = rdb.Do(context.Background(), "decr", courseCapacityKey).Int64()
	if err != nil {
		return -1
	}
	return ret
}

func HaveCourse(courseCapacityKey string) (ret bool) {
	res := rdb.Exists(context.Background(), courseCapacityKey).Val()
	return res == 1
}

// HaveBookCourse 验证学生是否抢到课程
func HaveBookCourse(bookCourseKey, value string) bool {
	val := rdb.SIsMember(context.Background(), bookCourseKey, value).Val()
	return val
}

// BookCourseInRedis 学生抢课，首先将抢课信息存入redis中
func BookCourseInRedis(bookCourseKey, value string) bool {
	val := rdb.SAdd(context.Background(), bookCourseKey, value).Val()
	return val == 1
}

func GetStudentCourse(studentId int64) ([]types.TCourse, e.ErrNo) {
	var courseList []types.TCourse
	err := db.Debug().Model(&types.TCourse{}).
		Select("course.course_id ,`name`,teacher_id").
		Joins("join course_book on (course.course_id = course_book.course_id)	").
		Where("course_book.student_id = ?", studentId).
		Scan(&courseList).Error
	if err != nil {
		return nil, e.UnknownError
	}
	return courseList, e.OK
}

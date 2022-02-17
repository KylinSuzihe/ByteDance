package student_service

import (
	"demo1/common"
	"demo1/models"
	"demo1/pkg/e"
	"demo1/pkg/types"
	"fmt"
	"strconv"
)

// BookCourse 使用redis+阻塞队列实现
func BookCourse(studentId, courseId string) e.ErrNo {
	// 用来缓存课程数量的key
	courseCapacityKey := "course:" + courseId
	// 用来缓存学生抢课信息的key
	bookCourseKey := "student:" + studentId

	// 首先从redis中查看是否存在课程
	if !models.HaveCourse(courseCapacityKey) {
		return e.CourseNotExisted
	}

	// 从redis中查询是否抢过课程，如果已经抢过课程，返回课程不可获取
	if models.HaveBookCourse(bookCourseKey, courseId) {
		return e.CourseNotAvailable
	}

	// 从redis中查询课程是否还有容量，如果没有直接返回
	if !models.HaveRemainderCapacity(courseCapacityKey) {
		return e.CourseNotAvailable
	}

	// 如果减完课程之后小于0，返回
	remain := models.DecrCapacity(courseCapacityKey)
	if remain < 0 {
		return e.CourseNotAvailable
	}

	// 向redis中存入抢课信息
	if !models.BookCourseInRedis(bookCourseKey, courseId) {
		return e.CourseNotAvailable
	}
	msg := [2]string{studentId, courseId}
	channel := common.GetChannel(0)
	*channel <- msg
	return e.OK
}

func ChannelConsumer() {
	for {
		msg, ok := <-(*common.GetChannel(0))
		if !ok {
			continue
		}
		errno := models.BookCourse(msg[0], msg[1])
		fmt.Println(e.GetMsg(errno))
	}
}

func GetStudentCourse(studentId string) ([]types.TCourse, e.ErrNo) {
	studentId1, err := strconv.ParseInt(studentId, 10, 64)
	if err != nil {
		return nil, e.ParamInvalid
	}

	// 需要验证学生是否存在
	if !models.ExistMemberByUserId(studentId1, types.Student) {
		return nil, e.StudentNotExisted
	}

	courseList, errno := models.GetStudentCourse(studentId1)
	return courseList, errno
}

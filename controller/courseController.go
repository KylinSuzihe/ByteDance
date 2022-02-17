package controller

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
	"demo1/service/course_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 作者：田思润

func CreateCourse(r *gin.Context) {
	var createCourseRequest types.CreateCourseRequest
	if err := r.ShouldBindJSON(&createCourseRequest); err != nil {
		r.JSON(http.StatusBadRequest, types.CreateCourseResponse{
			Code: e.ParamInvalid,
			Data: struct{ CourseID string }{CourseID: ""},
		})
		return
	}

	courseId, errno := course_service.CreateCourse(createCourseRequest.Name, createCourseRequest.Cap)

	r.JSON(http.StatusOK, types.CreateCourseResponse{
		Code: errno,
		Data: struct{ CourseID string }{CourseID: strconv.FormatInt(courseId, 10)},
	})
	return
}

func GetCourse(r *gin.Context) {
	var getCourseRequest types.GetCourseRequest
	var exist bool
	getCourseRequest.CourseID, exist = r.GetQuery("CourseID")

	// 如果参数不存在
	if !exist {
		r.JSON(http.StatusBadRequest, types.GetCourseResponse{
			Code: e.ParamInvalid,
			Data: struct {
				CourseID  string
				Name      string
				TeacherID string
			}{CourseID: "", Name: "", TeacherID: ""},
		})
		return
	}

	tCourse, errno := course_service.GetCourse(getCourseRequest.CourseID)
	if tCourse == nil {
		r.JSON(http.StatusOK, types.GetCourseResponse{
			Code: errno,
			Data: struct {
				CourseID  string
				Name      string
				TeacherID string
			}{CourseID: "", Name: "", TeacherID: ""},
		})
	} else {
		r.JSON(http.StatusOK, types.GetCourseResponse{
			Code: errno,
			Data: *tCourse,
		})
	}
	return
}

func ScheduleCourse(r *gin.Context) {
	var scheduleCourseRequest types.ScheduleCourseRequest
	if err := r.BindJSON(&scheduleCourseRequest); err != nil {
		r.JSON(http.StatusBadRequest, types.ScheduleCourseResponse{
			Code: e.ParamInvalid,
			Data: nil,
		})
		return
	}
	scheduleCourse, errNo := course_service.ScheduleCourse(scheduleCourseRequest.TeacherCourseRelationShip)

	r.JSON(http.StatusOK, types.ScheduleCourseResponse{
		Code: errNo,
		Data: scheduleCourse,
	})
	return
}

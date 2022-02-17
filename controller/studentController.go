package controller

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
	"demo1/service/student_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BookCourse(context *gin.Context) {
	var bookCourseRequest types.BookCourseRequest
	if err := context.ShouldBindJSON(&bookCourseRequest); err != nil {
		context.JSON(http.StatusBadRequest, types.BookCourseResponse{
			Code: e.ParamInvalid,
		})
		return
	}
	errno := student_service.BookCourse(bookCourseRequest.StudentID, bookCourseRequest.CourseID)
	context.JSON(http.StatusOK, types.BookCourseResponse{
		Code: errno,
	})
	return
}

func GetStudentCourse(context *gin.Context) {
	var getStudentCourseRequest types.GetStudentCourseRequest
	var exist bool
	getStudentCourseRequest.StudentID, exist = context.GetQuery("StudentID")
	if !exist {
		context.JSON(http.StatusBadRequest, types.GetStudentCourseResponse{
			Code: e.ParamInvalid,
			Data: struct{ CourseList []types.TCourse }{CourseList: nil},
		})
		return
	}
	courseList, errno := student_service.GetStudentCourse(getStudentCourseRequest.StudentID)
	context.JSON(http.StatusOK, types.GetStudentCourseResponse{
		Code: errno,
		Data: struct{ CourseList []types.TCourse }{CourseList: courseList},
	})
	return
}

package controller

import (
	"demo1/pkg/e"
	"demo1/pkg/types"
	"demo1/service/teacher_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 作者：田思润

func BindCourse(r *gin.Context) {
	var bindCourseRequest types.BindCourseRequest
	if err := r.BindJSON(&bindCourseRequest); err != nil {
		r.JSON(http.StatusBadRequest, types.BindCourseResponse{
			Code: e.ParamInvalid,
		})
		return
	}
	errno := teacher_service.BindCourse(bindCourseRequest.CourseID, bindCourseRequest.TeacherID)

	r.JSON(http.StatusOK, types.BindCourseResponse{
		Code: errno,
	})
	return
}

func UnbindCourse(r *gin.Context) {
	var unbindCourseRequest types.UnbindCourseRequest
	if err := r.BindJSON(&unbindCourseRequest); err != nil {
		r.JSON(http.StatusBadRequest, types.UnbindCourseResponse{
			Code: e.ParamInvalid,
		})
		return
	}
	errno := teacher_service.UnbindCourse(unbindCourseRequest.CourseID, unbindCourseRequest.TeacherID)
	r.JSON(http.StatusOK, types.UnbindCourseResponse{
		Code: errno,
	})
	return
}

func GetTeacherCourse(r *gin.Context) {
	var getTeacherCourseRequest types.GetTeacherCourseRequest
	var exist bool
	getTeacherCourseRequest.TeacherID, exist = r.GetQuery("TeacherID")
	if !exist {
		r.JSON(http.StatusBadRequest, types.GetTeacherCourseResponse{
			Code: e.ParamInvalid,
			Data: struct{ CourseList []*types.TCourse }{CourseList: nil},
		})
		return
	}
	courseList, errno := teacher_service.GetTeacherCourse(getTeacherCourseRequest.TeacherID)
	r.JSON(http.StatusOK, types.GetTeacherCourseResponse{
		Code: errno,
		Data: struct{ CourseList []*types.TCourse }{CourseList: courseList},
	})
	return
}

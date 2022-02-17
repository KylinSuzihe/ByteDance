package routers

import (
	"demo1/controller"
	"demo1/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	//路由上加入session中间件
	router.Use(sessions.Sessions("camp-session", store))

	g := router.Group("/api/v1")
	// 登录、权限验证
	g.POST("/auth/login", controller.Login)
	g.POST("/auth/logout", controller.Logout)
	g.GET("/auth/whoami", controller.WhoAmI)

	// 成员管理
	g.POST("/member/create", controller.CreateMember)
	g.GET("/member/", controller.GetMember)
	g.GET("/member/list", controller.GetMemberList)
	g.POST("/member/update", controller.UpdateMember)
	g.POST("/member/delete", controller.DeleteMember)

	// 排课
	g.POST("/course/create", controller.CreateCourse)
	g.GET("/course/get", controller.GetCourse)
	g.POST("/course/schedule", controller.ScheduleCourse)

	g.POST("/teacher/bind_course", controller.BindCourse)
	g.POST("/teacher/unbind_course", controller.UnbindCourse)
	g.GET("/teacher/get_course", controller.GetTeacherCourse)

	// 抢课
	g.POST("/student/book_course", controller.BookCourse)
	g.GET("/student/course", controller.GetStudentCourse)

	return router
}

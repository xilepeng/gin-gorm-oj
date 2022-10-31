package router

import (
	"gin-gorm-oj/service"

	_ "gin-gorm-oj/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 路由规则
	// r.GET("/ping", service.Ping)

	// swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 问题
	r.GET("/problem-list", service.GetProblemList) // http://localhost:8080/problem-list
	r.GET("/problem-detail", service.GetProblemDetail)
	// url 后多输入1个空格导致：404 page not found

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	// 排行榜
	r.GET("/rank-list", service.GetRankList)
	// 提交记录
	r.GET("/submit-list", service.GetSubmitList)

	return r
}

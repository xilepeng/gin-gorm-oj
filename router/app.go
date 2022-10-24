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
	r.GET("/ping", service.Ping)
	// url 后多输入1个空格导致：404 page not found
	r.GET("/problem-list", service.GetProblemList) // http://localhost:8080/problem-list

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
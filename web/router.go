package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rodert/hepburn/internal/handler"
	"github.com/rodert/hepburn/service/hello"
)

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) // 设置 release模式
	gin.SetMode(gin.DebugMode) // 设置为 debug 模式
	router := gin.New()
	router.Use(handler.Cors()) // 允许跨域

	service := router.Group("hello")
	{
		service.GET("hello", handler.TRPathParamHandler(hello.Hello))
	}
	return router
}

package router

import (
	"github.com/gin-gonic/gin"
	my "sam.com/go-api/api/helloworld"
)

func InitRouter() *gin.Engine {
	// 建立路由
	router := gin.Default()

	// v1 := router.Group("/Demo")
	router.GET("/demo/v1/hello", my.Hello)
	router.POST("/demo/v1/hi", my.Hi)

	person := router.Group("/user/person")
	{
		person.POST("signup", my.Signup)
		person.POST("login", my.Signup)
	}
	// router.GET("/users", Users)

	// router.POST("/user", Store)

	// router.PUT("/user/:id", Update)

	// router.DELETE("/user/:id", Destroy)

	return router
}

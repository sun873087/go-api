package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"sam.com/go-api/api/router"
	_ "sam.com/go-api/docs"
)

// @title Sam Gin Swagger Test
// @version 1.0.0
// @description 給Flutter學習使用的API
// @host localhost:8080
// @contact.name sam.cheng
// @contact.email sam.cheng@wnc.com.tw
func main() {
	router := router.InitRouter()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run()
}

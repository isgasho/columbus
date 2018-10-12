package main

import (
	"columbus/utils"
	_ "columbus/docs"
	"columbus/api"
	//"columbus/config"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1/

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	// 注册
	accountsRegister(v1)
	v1.Use(utils.JWTAuth()) // 使用中间件 进行auth校验
    accountsManager(v1)

    // 丝袜哥
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func accountsRegister(r *gin.RouterGroup) {
	c := api.NewController()
	r.POST("/accounts/token", c.AccountToken)
}

func accountsManager(r *gin.RouterGroup) {
	c := api.NewController()
	r.GET(":id", c.ShowAccount)
	r.GET("", c.ListAccounts)
	r.POST("", c.AddAccount)
	r.DELETE(":id", c.DeleteAccount)
	r.PATCH(":id", c.UpdateAccount)
}
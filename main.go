package main

import (
	"columbus/utils"
	_ "columbus/docs"
	"columbus/api"
	_ "columbus/database"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1/

// @securityDefinitions.apiKey Bearer
// @type apiKey
// @in header
// @name Authorization

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	// 注册
	c := api.NewController()
	accountsRegister(v1, c)
	v1.Use(utils.JWTAuth()) // 使用中间件 进行auth校验
	{
        accountsManager(v1, c)
    }

    // 丝袜哥
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func accountsRegister(r *gin.RouterGroup, c *api.Controller) {
	r.POST("/accounts/token", c.AccountToken)
}

func accountsManager(r *gin.RouterGroup, c *api.Controller) {
	r.GET("/accounts/:id", c.ShowAccount)
	r.GET("/accounts", c.ListAccounts)
	r.POST("/accounts", c.AddAccount)
	r.DELETE("/accounts/:id", c.DeleteAccount)
	r.PATCH("/accounts/:id", c.UpdateAccount)
}
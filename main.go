package main

import (
	//"fmt"
	"errors"
	"net/http"

	"columbus/httputil"
	_ "columbus/docs"
	"columbus/controller"
	"columbus/config"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "columbus config file path.")
)

func main() {
	r := gin.Default()

	c := controller.NewController()

	// 读取配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// TODO 更优雅的router方式
	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}

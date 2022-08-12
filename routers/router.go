package routers

import (
	v1 "ums/api/v1"
	"ums/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	userRouter := r.Group("/api/v1")
	{
		// 用户管理
		userRouter.POST("user", v1.AddUser)
		userRouter.PUT("user", v1.UpdateUser)
		userRouter.POST("user/del", v1.DeleteUser)
		userRouter.GET("user", v1.PageUser)
	}

	r.Run(":8006")
}

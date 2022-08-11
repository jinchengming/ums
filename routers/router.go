package routers

import (
	v1 "ums/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	authRouter := r.Group("/api/v1")
	{
		// 用户管理
		authRouter.POST("user", v1.AddUser)
		authRouter.PUT("user", v1.UpdateUser)
		authRouter.POST("user/del", v1.DeleteUser)
		authRouter.GET("user", v1.PageUser)
	}

	r.Run(":8006")
}

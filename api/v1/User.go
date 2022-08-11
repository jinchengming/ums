package v1

import (
	"net/http"
	"ums/model"
	"ums/utils"
	"ums/utils/result"

	"github.com/gin-gonic/gin"
)

var code int

func AddUser(c *gin.Context) {
	code = result.SUCCESS
	var user model.User
	c.ShouldBindJSON(&user)
	exist := model.CheckUserExist(user.NickName)
	if exist {
		// 已存在 抛出错误信息
		code = result.UsernameUsed
	} else {
		code = model.CreateUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": result.GetErrMsg(code),
	})
}

func UpdateUser(c *gin.Context) {
	code = result.SUCCESS
	var user model.User
	c.ShouldBindJSON(&user)
	oldUser := model.GetUser(user.ID)
	if oldUser.NickName != user.NickName {
		exist := model.CheckUserExist(user.NickName)
		if exist {
			// 已存在 抛出错误信息
			code = result.UsernameUsed
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": result.GetErrMsg(code),
			})
			return
		}
	}
	code = model.ModifyUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": result.GetErrMsg(code),
	})
}

func DeleteUser(c *gin.Context) {
	var ids []int
	c.ShouldBindJSON(&ids)
	code := model.RemoveUser(ids)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": result.GetErrMsg(code),
	})
}

func PageUser(c *gin.Context) {
	code = result.SUCCESS
	pageNum, pageSize := utils.GetPageParam(c)
	name := c.DefaultQuery("name", "")
	userList, total := model.PageUser(pageNum, pageSize, name)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": result.GetErrMsg(code),
		"data":    userList,
		"total":   total,
	})
}

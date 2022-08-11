package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageParam(c *gin.Context) (pageNum, pageSize int) {
	pageNum, _ = strconv.Atoi(c.DefaultQuery("pageNum", "20"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	return
}

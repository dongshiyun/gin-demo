package controllers

import (
	"gd/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	rs, err := models.GetUserInfo(userId)
	if err == nil {
		ReturnSuccess(c, 0, "success", rs, 1)
		return
	}
	ReturnError(c, 4004, "没有相关信息")
}

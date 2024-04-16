package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonStruct{Code: code, Msg: msg}
	c.JSON(200, json)
}

//格式化时间
func DateFormat(times int64) string {
	video_time := time.Unix(times, 0)
	return video_time.Format("2006-01-02")
}

//获取页码信息
func PageFormat(c *gin.Context) (int, int) {
	pageStr := c.DefaultPostForm("page", "")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := c.DefaultPostForm("pageSize", "")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	limit := pageSize

	return offset, limit
}

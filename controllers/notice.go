package controllers

import (
	"gd/models"

	"github.com/gin-gonic/gin"
)

type NoticeData struct {
	Id               int    `json:"id"`
	UserId           int    `json:"userId"`
	SendUserId       int    `json:"sendUserId"`
	SendUserNickName string `json:"sendUserNickName"`
	SendUserAvatar   string `json:"sendUserAvatar"`
	Content          string `json:"content"`
	AddTimeStr       string `json:"addTimeStr"`
}

func GetNotices(c *gin.Context) {
	type PostData struct {
		UserId int `json:"userId"`
	}
	json := PostData{}
	c.BindJSON(&json)
	userId := json.UserId

	rs, err := models.GetNotices(userId)
	if err == nil {
		count, _ := models.GetNoticesCount(userId)

		var data []NoticeData
		var item NoticeData
		for _, v := range rs {
			item.Id = v.Id
			item.UserId = v.UserId
			item.SendUserId = v.SendUserId
			item.Content = v.Content
			if v.SendUserId == 0 {
				item.SendUserNickName = "系统消息"
				item.SendUserAvatar = "/homestatic/img/avatar.jpeg"
			} else {
				//调用用户服务
				item.SendUserNickName = "系统消息"
				item.SendUserAvatar = "/homestatic/img/avatar.jpeg"
			}
			item.AddTimeStr = DateFormat(v.AddTime)
			data = append(data, item)
		}

		ReturnSuccess(c, 0, "success", data, count)
		return
	}
	ReturnError(c, 4004, "没有相关信息")
}

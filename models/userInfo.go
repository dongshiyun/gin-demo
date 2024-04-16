package models

import (
	"gd/initDb"
)

type UserInfo struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Sex          string `json:"sex"`
	Introduction string `json:"introduction"`
	Homepage     string `json:"homepage"`
	AddTime      int64  `json:"addTime"`
	UpdateTime   string `json:"updateTime"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

func GetUserInfo(userId int) (UserInfo, error) {
	var userInfo UserInfo
	db := initDb.Db()
	rs := db.Where("user_id = ?", userId).First(&userInfo)
	return userInfo, rs.Error
}

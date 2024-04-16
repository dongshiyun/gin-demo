package models

import (
	"mjs/initDb"
)

type Notice struct {
	Id         int
	UserId     int
	SendUserId int
	Content    string
	AddTime    int64
}

func (Notice) TableName() string {
	return "notice"
}

func GetNotices(userId int) ([]Notice, error) {
	var notices []Notice
	db := initDb.Db()
	rs := db.Where("user_id = ?", userId).Order("add_time desc, id desc").Find(&notices)
	return notices, rs.Error
}

func GetNoticesCount(userId int) (int, error) {
	var notices []Notice
	var count int = 0
	db := initDb.Db()
	rs := db.Where("user_id = ?", userId).Find(&notices).Count(&count)
	return count, rs.Error
}

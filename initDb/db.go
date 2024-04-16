package initDb

import (
	"gd/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Db() *gorm.DB {
	db, err := gorm.Open("mysql", config.Mysqldb)
	//defer db.Close()
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	return db
}

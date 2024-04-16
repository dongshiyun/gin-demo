package main

import (
	"gd/router"
)

func main() {
	r := router.Router()

	//consul.Register("127.0.0.1", 8801, "goNoticeHttp", []string{"go"}, "goNoticeHttp")

	//consul.GetServiceByID("articleHttp")

	r.Run(":8801") // listen and serve on 0.0.0.0:8080
}

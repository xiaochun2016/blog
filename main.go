package main

import (
	"blog/db"
	"blog/route"
)

func main() {

	err := db.Init()
	if err != nil {
		panic("初始化数据库失败")
	}
	r := route.RegisterRoute()
	_ = r.Run(":8080")
}

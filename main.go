package main

import (
	"blog/controller"
	"blog/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := db.Init("root:123456@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		fmt.Println(err)
		return
	}

	v1 := r.Group("/book")
	{
		v1.GET("/", controller.Index)      //主页
		v1.POST("/add", controller.AddArt) //添加文章
	}
	cateV1 := r.Group("/cate")
	{
		cateV1.POST("/add", controller.AddCate)         //添加分类
		cateV1.GET("/all", controller.GetAll)           //全部分类
		cateV1.POST("/getById", controller.GetCateById) //通过id获取分类
	}

	_ = r.Run(":8080")
}

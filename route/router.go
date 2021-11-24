package route

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/article")
	{
		v1.GET("/", controller.Index)                //主页
		v1.POST("/add", controller.AddArt)           //添加文章
		v1.POST("/list", controller.ArtList)         //文章列表
		v1.POST("/listById", controller.ArtListById) //通过分类id获取文章列表
	}
	cateV1 := r.Group("/cate")
	{
		cateV1.POST("/add", controller.AddCate)         //添加分类
		cateV1.GET("/all", controller.GetAll)           //全部分类
		cateV1.POST("/getById", controller.GetCateById) //通过id获取分类
	}
	leave := r.Group("leave")
	{
		leave.POST("add", controller.LeaveAdd) //添加留言
	}
	return r
}

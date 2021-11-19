package controller

import (
	"blog/db"
	"blog/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": "",
	})
}

//AddArt 添加文章
func AddArt(c *gin.Context) {
	title := c.PostForm("title")
	cateIdStr := c.PostForm("cate_id")
	cateId, _ := strconv.ParseInt(cateIdStr, 10, 10)
	summary := c.PostForm("summary")
	email := c.PostForm("email")
	content := c.PostForm("content")
	username := c.PostForm("name")
	if title == "" || cateId < 1 || summary == "" || email == "" || content == "" || username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  "参数错误",
		})
		return
	}
	id, err := db.AddArticle(username, title, email, content, cateId, summary)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "添加失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "添加成功",
			"data": map[string]int64{"id": id},
		})
		return
	}
}

// ArtList 文件列表
func ArtList(c *gin.Context) {
	var param *model.PostParams
	err := c.ShouldBindWith(&param, binding.Form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  err.Error(),
		})
		return
	}
	list, err := db.ArtList(param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  err.Error(),
		})
		return
	}
	title := "获取成功"
	if len(list) < 1 {
		title = "暂无数据"
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  title,
			"data": list,
		})
		return
	}
	var ids []int64
	for _, v := range list {
		ids = append(ids, int64(v.CateId))
	}
	idList, _ := db.GetCateByIds(ids)
	for _, v := range list {
		val, ok := idList[int64(v.CateId)]
		if ok {
			v.CateName = val
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  title,
		"data": list,
	})
	return
}

//ArtListById 根据分类获取文章列表
func ArtListById(c *gin.Context) {
	var param *model.PostParams
	err := c.ShouldBindWith(&param, binding.Form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  err.Error(),
		})
		return
	}
	if param.CateId < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  "分类id不能为空",
		})
		return
	}
	list, err := db.ArtListById(param.CateId, param.Page, param.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  err.Error(),
		})
		return
	}
	title := "获取成功"
	if len(list) < 1 {
		title = "暂无数据"
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  title,
			"data": list,
		})
		return
	}
	var ids []int64
	for _, v := range list {
		ids = append(ids, int64(v.CateId))
	}
	idList, _ := db.GetCateByIds(ids)
	for _, v := range list {
		val, ok := idList[int64(v.CateId)]
		if ok {
			v.CateName = val
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  title,
		"data": list,
	})
	return
}

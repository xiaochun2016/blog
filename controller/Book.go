package controller

import (
	"blog/db"
	"github.com/gin-gonic/gin"
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

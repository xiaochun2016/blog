package controller

import (
	"blog/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//AddCate 添加分类
func AddCate(c *gin.Context) {
	cateName := c.PostForm("cate_name")
	cateSort := c.PostForm("cate_sort")
	cateSortNum, _ := strconv.ParseInt(cateSort, 10, 10)
	if cateName == "" || cateSortNum < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": "202",
			"msg":  "参数错误",
		})
		return
	}
	id, err := db.AddCate(cateName, cateSortNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "添加失败",
			"date": err.Error(),
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

//GetAll 获取全部分类
func GetAll(c *gin.Context) {
	cate, err := db.AllCate()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "获取失败",
			"date": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "获取成功",
			"date": cate,
		})
	}
}
func GetCateById(c *gin.Context) {
	idsStr := c.PostForm("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "参数错误",
		})
		return
	}
	ids := strings.Split(idsStr, ",")
	if len(ids) < 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "参数错误",
		})
		return
	}
	idsNum := make([]int64, 0)
	for _, v := range ids {
		num, _ := strconv.Atoi(v)
		if num > 0 {
			idsNum = append(idsNum, int64(num))
		}
	}
	cate, err := db.GetCateByIds(idsNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "获取失败",
			"date": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "201",
			"msg":  "获取成功",
			"date": cate,
		})
	}
}

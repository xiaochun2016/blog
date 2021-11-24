package controller

import (
	"blog/db"
	"blog/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

//留言

// LeaveAdd 添加留言
func LeaveAdd(c *gin.Context) {
	var LeaveAddInfo *model.AddLeave
	err := c.ShouldBindWith(&LeaveAddInfo, binding.Form)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 202,
			"msg":  err.Error(),
		})
		return
	}
	id, err := db.AddLeave(LeaveAddInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
		"data": id,
	})
	return
}

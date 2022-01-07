package api

import (
	"TodoList/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建
func CreateTask(c *gin.Context) {
	var createService service.CreateTaskService
	if err := c.ShouldBind(&createService); err == nil {
		uid, _ := c.Get("uid")
		// 新增
		res := createService.Create(uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 展示某条数据
func ShowTask(c *gin.Context) {
	var showService service.ShowTaskService
	uid, _ := c.Get("uid")
	res := showService.Show(uid.(uint), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// 查询所有数据
func ListTask(c *gin.Context) {
	var listService service.ListTaskService
	if err := c.ShouldBind(&listService); err == nil {
		uid, _ := c.Get("uid")
		res := listService.List(uid.(uint))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 修改
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err == nil {
		uid, _ := c.Get("uid")
		tid := c.Param("id")
		res := updateTask.Update(uid.(uint), tid)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// 删除
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	uid, _ := c.Get("uid")
	res := deleteTask.Delete(uid.(uint), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

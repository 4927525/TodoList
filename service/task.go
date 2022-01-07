package service

import (
	"TodoList/model"
	"TodoList/serializer"
	"net/http"
	"time"
)

type CreateTaskService struct {
	Uid       uint   `json:"uid"`
	Title     string `json:"title" form:"title" binding:"required"`
	Status    uint   `json:"status" form:"status"`
	Content   string `json:"content" form:"content"`
	StartTime int64
	EndTime   int64
}

func (service *CreateTaskService) Create(uid uint) serializer.Response {
	var task model.Task
	task.Uid = uid
	task.Title = service.Title
	task.Status = service.Status
	task.Content = service.Content
	task.StartTime = time.Now().Unix()
	task.EndTime = 0
	if err := model.Db.Create(&task).Error; err != nil {
		return serializer.Response{
			Status: http.StatusOK,
			Msg:    "新增失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "新增成功",
	}
}

type ShowTaskService struct {
}

// 展示某条数据
func (service *ShowTaskService) Show(uid uint, id string) serializer.Response {
	var task model.Task
	count := 0
	err := model.Db.Model(&model.Task{}).Where("uid = ? And id = ?", uid, id).First(&task).Count(&count).Error
	if err != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Data:   serializer.BuildTask(task),
		Msg:    "查询成功",
	}
}

type ListTaskService struct {
	Search   string `json:"search" form:"search"`
	PageSize int    `json:"page_size" form:"page_size"`
	PageNum  int    `json:"page_num" form:"page_num"`
}

// 查询所有数据
func (service *ListTaskService) List(uid uint) serializer.Response {
	var task []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	offset := (service.PageNum - 1) * service.PageSize
	err := model.Db.Model(&model.Task{}).
		Where("uid = ?", uid).
		Where("title like ? Or content like ?", "%"+service.Search+"%", "%"+service.Search+"%").
		Limit(service.PageSize).Offset(offset).
		Count(&count).Find(&task).Error
	if err != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "查询失败",
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(task), count)
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Status  uint   `json:"status" form:"status"`
	Content string `json:"content" form:"content"`
	EndTime int64
}

// 修改
func (service *UpdateTaskService) Update(uid uint, id string) serializer.Response {
	var task model.Task
	model.Db.Model(&model.Task{}).Where("uid = ?", uid).First(&task, id)
	task.Title = service.Title
	task.Status = service.Status
	task.Content = service.Content
	task.EndTime = service.EndTime
	err := model.Db.Model(&model.Task{}).Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status: http.StatusInternalServerError,
			Msg:    "修改失败",
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "修改成功",
		Data:   serializer.BuildTask(task),
	}
}

type DeleteTaskService struct {
}

// 删除
func (service *DeleteTaskService) Delete(uid uint, tid string) serializer.Response {
	var task model.Task
	model.Db.Model(&model.Task{}).Where("uid = ?", uid).Delete(&task, tid)
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    "删除成功",
	}
}

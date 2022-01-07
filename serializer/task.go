package serializer

import (
	"TodoList/model"
	"time"
)

type Task struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Status    uint   `json:"status"`
	Content   string `json:"content"`
	StartTime string `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// 构建单条
func BuildTask(task model.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		StartTime: time.Unix(task.StartTime, 0).Format("2006-01-02 15:04:05"),
		EndTime:   task.EndTime,
	}
}

// 构建多条
func BuildTasks(items []model.Task) (tasks []interface{}) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}

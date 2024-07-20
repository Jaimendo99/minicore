package models

import (
	"gorm.io/gorm"
	"time"
)

// Task model
type Task struct {
	gorm.Model
	Description   string
	StartDateTS   int32
	EndDateTS     int32
	EstimatedDays int32
	Status        string
	EmployeeID    uint
	ProjectID     uint
	Employee      Employee `gorm:"foreignKey:EmployeeID"`
	Project       Project  `gorm:"foreignKey:ProjectID"`
}

type TaskUI struct {
	ID            int
	Description   string
	StartDate     string
	EndDate       string
	EstimatedDays int32
	Status        string
}

// TaskToTaskUI converts a Task to a TaskUI
func TaskToTaskUI(task Task) TaskUI {
	var startDate, endDate string
	if task.StartDateTS != 0 {
		startDate = int32ToDateString(task.StartDateTS)
	}
	if task.EndDateTS != 0 {
		endDate = int32ToDateString(task.EndDateTS)
	}
	return TaskUI{
		ID:            int(task.ID),
		Description:   task.Description,
		StartDate:     startDate,
		EndDate:       endDate,
		EstimatedDays: task.EstimatedDays,
		Status:        task.Status,
	}
}

// TaskToTaskUIList converts a list of Task to a list of TaskUI
func TaskToTaskUIList(tasks []Task) []TaskUI {
	var taskUIList []TaskUI
	for _, task := range tasks {
		taskUIList = append(taskUIList, TaskToTaskUI(task))
	}
	return taskUIList

}

func int32ToDateString(ts int32) string {
	return time.Unix(int64(ts), 0).Format("2006-01-02")
}

// DateStringToInt32 convierte una cadena de fecha en formato "YYYY-MM-DD" a un timestamp int32
func DateStringToInt32(dateStr string) (int32, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, err
	}
	return int32(t.Unix()), nil
}

// FilterTasksByStartDate filtra las tareas por la fecha de inicio
func FilterTasksByStartDate(tasks []Task, startDateTS int32) []Task {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.StartDateTS >= startDateTS {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

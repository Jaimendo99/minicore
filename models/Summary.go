package models

import (
	"time"
)

type OverDue struct {
	EmployeeName string
	TaskName     string
	ProjectName  string
	StartDate    string
	DueDate      string
	DueDays      int
}

func TaskToOverDue(task Task, refDate time.Time) (OverDue, error) {
	startDate := time.Unix(int64(task.StartDateTS), 0)
	dueDate := time.Unix(int64(task.EndDateTS), 0)
	dueDays := countBusinessDays(dueDate, refDate)

	employeeName := task.Employee.FirstName + " " + task.Employee.LastName
	projectName := task.Project.Name

	return OverDue{
		EmployeeName: employeeName,
		TaskName:     task.Description,
		ProjectName:  projectName,
		StartDate:    startDate.Format("2006-01-02"),
		DueDate:      dueDate.Format("2006-01-02"),
		DueDays:      dueDays,
	}, nil
}

func countBusinessDays(start, end time.Time) int {
	businessDays := 0
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			businessDays++
		}
	}
	return businessDays
}

func TaskToOverDueList(tasks []Task, refDate time.Time) ([]OverDue, error) {
	var overDueList []OverDue
	for _, task := range tasks {
		overDue, err := TaskToOverDue(task, refDate)
		if err != nil {
			return nil, err
		}
		overDueList = append(overDueList, overDue)
	}
	return overDueList, nil
}

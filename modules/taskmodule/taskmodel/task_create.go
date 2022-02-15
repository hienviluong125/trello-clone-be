package taskmodel

import "errors"

type TaskCreate struct {
	Title        *string `json:"title" gorm:"column:title;"`
	Body         *string `json:"body" gorm:"column:body;"`
	Index        *int    `json:"index" gorm:"column:index;"`
	ListId       *int    `json:"list_id" gorm:"column:list_id;"`
	Status       *bool   `json:"-" gorm:"column:status;"`
	ReportedById *int    `json:"-" gorm:"column:reported_by_id;"`
	AssigneeId   *int    `json:"assignee_id" gorm:"column:assignee_id;"`
}

func (TaskCreate) TableName() string {
	return "tasks"
}

func (t *TaskCreate) Validate() error {
	if t.Title == nil || len(*t.Title) == 0 {
		return errors.New("title can't be blank")
	}

	if t.Index == nil {
		return errors.New("index can't be blank")
	}

	if t.ListId == nil {
		return errors.New("list id can't be blank")
	}

	if t.ReportedById == nil {
		return errors.New("repoted by id can't be blank")
	}

	// if t.AssigneeId == nil {
	// 	return errors.New("assignee id can't be blank")
	// }

	return nil
}

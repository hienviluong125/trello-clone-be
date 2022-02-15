package taskmodel

const Entity = "Task"

type Task struct {
	Id           int    `json:"id" gorm:"column:id;"`
	Title        string `json:"title" gorm:"column:title;"`
	Body         string `json:"body" gorm:"column:body;"`
	Index        int    `json:"index" gorm:"column:index;"`
	ListId       int    `json:"-" gorm:"column:list_id;"`
	Status       bool   `json:"-" gorm:"column:status;"`
	ReportedById int    `json:"-" gorm:"column:reported_by_id;"`
	AssigneeId   int    `json:"-" gorm:"column:assignee_id;"`
}

func (Task) TableName() string {
	return "tasks"
}

package taskmodel

type TaskUpdate struct {
	Title        *string `json:"title" gorm:"column:title;"`
	Body         *string `json:"body" gorm:"column:body;"`
	Index        *int    `json:"index" gorm:"column:index;"`
	ListId       *int    `json:"list_id" gorm:"column:list_id;"`
	Status       *bool   `json:"-" gorm:"column:status;"`
	ReportedById *int    `json:"-" gorm:"column:reported_by_id;"`
	AssigneeId   *int    `json:"assignee_id" gorm:"column:assignee_id;"`
}

func (TaskUpdate) TableName() string {
	return "tasks"
}

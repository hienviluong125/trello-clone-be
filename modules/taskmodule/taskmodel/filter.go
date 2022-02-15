package taskmodel

type Filter struct {
	Title      string `json:"title" gorm:"column:title;"`
	Body       string `json:"body" gorm:"column:body;"`
	Index      int    `json:"index" gorm:"column:index;"`
	ListId     int    `json:"list_id" gorm:"column:list_id;"`
	AssigneeId int    `json:"assignee_id" gorm:"column:assignee_id;"`
}

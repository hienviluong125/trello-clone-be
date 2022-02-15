package listmodel

import "hienviluong125/trello-clone-be/modules/taskmodule/taskmodel"

const Entity = "List"

type List struct {
	Id      int               `json:"id" gorm:"column:id;"`
	Name    string            `json:"name" gorm:"column:name;"`
	Index   int               `json:"index" gorm:"column:index;"`
	Status  bool              `json:"-" gorm:"column:status;"`
	BoardId int               `json:"-" gorm:"column:board_id;"`
	Tasks   []*taskmodel.Task `json:"tasks" gorm:"preload:false;"`
}

func (List) TableName() string {
	return "lists"
}

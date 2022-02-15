package listmodel

const Entity = "List"

type List struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Index   int    `json:"index" gorm:"column:index;"`
	Status  bool   `json:"-" gorm:"column:status;"`
	BoardId int    `json:"-" gorm:"column:board_id;"`
}

func (List) TableName() string {
	return "lists"
}

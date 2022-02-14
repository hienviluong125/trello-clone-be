package boardmodel

import "hienviluong125/trello-clone-be/modules/usermodule/usermodel"

const Entity = "Board"

type Board struct {
	Id      int               `json:"id" gorm:"column:id;"`
	Name    string            `json:"name" gorm:"column:name;"`
	Status  bool              `json:"-" gorm:"column:status;"`
	OwnerId int               `json:"-" gorm:"column:owner_id;"`
	Owner   *usermodel.User   `json:"owner,omitempty" gorm:"preload:false"`
	Members []*usermodel.User `json:"members,omitempty" gorm:"many2many:user_boards;preload:false"`
}

func (Board) TableName() string {
	return "boards"
}

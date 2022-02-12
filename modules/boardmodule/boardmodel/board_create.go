package boardmodel

import (
	"errors"
	"hienviluong125/trello-clone-be/modules/usermodule/usermodel"
)

type BoardCreate struct {
	Name    string `json:"name" gorm:"column:name;"`
	Status  bool   `json:"-" gorm:"column:status;"`
	OwnerId int    `json:"-" gorm:"column:owner_id;"`
	Owner   *usermodel.User
}

func (BoardCreate) TableName() string {
	return "boards"
}

func (board *BoardCreate) Validate() error {
	if len(board.Name) == 0 {
		return errors.New("name can't be blank")
	}

	return nil
}

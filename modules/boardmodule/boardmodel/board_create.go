package boardmodel

import (
	"errors"
	"hienviluong125/trello-clone-be/modules/usermodule/usermodel"

	"gorm.io/gorm"
)

type BoardCreate struct {
	Id      int    `json:"-" gorm:"column:id;"`
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

func (board *BoardCreate) AfterCreate(tx *gorm.DB) error {
	if board.OwnerId != 0 {
		var userBoard UserBoard = UserBoard{
			BoardId: board.Id,
			UserId:  board.OwnerId,
		}

		if err := tx.Create(&userBoard).Error; err != nil {
			return err
		}

		return nil
	}

	return errors.New("cannot create a board")
}

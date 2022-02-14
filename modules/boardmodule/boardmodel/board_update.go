package boardmodel

import (
	"errors"
)

type BoardUpdate struct {
	Name   string `json:"name" gorm:"column:name;"`
	Status *bool  `json:"-" gorm:"column:status;"`
}

func (BoardUpdate) TableName() string {
	return "boards"
}

func (board *BoardUpdate) Validate() error {
	if len(board.Name) == 0 {
		return errors.New("name can't be blank")
	}

	return nil
}

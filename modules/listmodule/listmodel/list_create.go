package listmodel

import "errors"

type ListCreate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Index   *int    `json:"index" gorm:"column:index;"`
	Status  *bool   `json:"-" gorm:"column:status;"`
	BoardId *int    `json:"-" gorm:"column:board_id;"`
}

func (ListCreate) TableName() string {
	return "lists"
}

func (list *ListCreate) Validate() error {
	if list.Name == nil || len(*list.Name) < 1 {
		return errors.New("name can't be blank")
	}

	if list.Index == nil {
		return errors.New("index can't be blank")
	}

	if list.BoardId == nil || *list.BoardId == 0 {
		return errors.New("board id can't be blank")
	}

	return nil
}

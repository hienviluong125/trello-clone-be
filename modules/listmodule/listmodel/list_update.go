package listmodel

import "errors"

type ListUpdate struct {
	Name   *string `json:"name" gorm:"column:name;"`
	Index  *int    `json:"index" gorm:"column:index;"`
	Status *bool   `json:"-" gorm:"column:status;"`
}

func (ListUpdate) TableName() string {
	return "lists"
}

func (list *ListUpdate) Validate() error {
	if list.Name == nil || len(*list.Name) < 1 {
		return errors.New("name can't be blank")
	}

	return nil
}

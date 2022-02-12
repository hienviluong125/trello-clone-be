package usermodel

type UserUpdate struct {
	Name         string `json:"name" gorm:"column:name;"`
	RefreshToken string `json:"-" gorm:"column:refresh_token"`
	Status       bool   `json:"-" gorm:"column:status"`
}

func (UserUpdate) TableName() string {
	return "users"
}

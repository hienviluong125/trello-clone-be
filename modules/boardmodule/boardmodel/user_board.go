package boardmodel

type UserBoard struct {
	UserId  int `json:"user_id" gorm:"column:user_id;"`
	BoardId int `json:"-" gorm:"column:board_id;"`
}

func (UserBoard) TableName() string {
	return "user_boards"
}

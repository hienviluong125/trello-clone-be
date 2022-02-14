package usermodel

const Entity = "User"

type User struct {
	Id             int    `json:"id" gorm:"column:id;"`
	Name           string `json:"name" gorm:"column:name;"`
	Email          string `json:"email" gorm:"column:email;unique"`
	HashedPassword string `json:"-" gorm:"column:hashed_password;"`
	Status         bool   `json:"-" gorm:"column:status"`
	Role           string `json:"-" gorm:"column:role"`
	RefreshToken   string `json:"-" gorm:"column:refresh_token"`
	// Boards         []*boardmodel.Board `json:"boards"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

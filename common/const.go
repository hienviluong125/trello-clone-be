package common

const CurrentUser = "CURRENT_USER"

const RefreshToken = "REFRESH_TOKEN"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

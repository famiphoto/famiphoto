package entities

type User struct {
	UserID  int64
	MyID    string
	Status  UserStatus
	IsAdmin bool
}

func NewInitUser(myID string, isAdmin bool) *User {
	return &User{
		UserID:  0,
		MyID:    myID,
		Status:  UserStatusEnabled,
		IsAdmin: isAdmin,
	}
}

type UserStatus int

const (
	UserStatusDisabled UserStatus = 0
	UserStatusEnabled  UserStatus = 1
)

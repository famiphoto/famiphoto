package entities

type User struct {
	UserID  string
	Status  UserStatus
	IsAdmin bool
}

func NewInitUser(userID string, isAdmin bool) *User {
	return &User{
		UserID:  userID,
		Status:  UserStatusEnabled,
		IsAdmin: isAdmin,
	}
}

type UserStatus int

const (
	UserStatusDisabled UserStatus = 0
	UserStatusEnabled  UserStatus = 1
)

package responses

import "github.com/famiphoto/famiphoto/api/entities"

type SignUpResponse struct {
	MyID string `json:"myId"`
}

func NewSignUpResponse(user *entities.User) *SignUpResponse {
	return &SignUpResponse{MyID: user.MyID}
}

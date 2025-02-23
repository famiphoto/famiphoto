package requests

import "github.com/labstack/echo/v4"

type SignUpRequest struct {
	MyID     string `json:"myId" validate:"required"`
	Password string `json:"password" validate:"required"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (r *SignUpRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.Validate(r)
}

type SignInRequest struct {
	MyID     string `json:"myId" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *SignInRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.Validate(r)
}

package rest

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Login handler authenticates user and returns a JWT Token
func (h *Handler) Login(c echo.Context) error {
	var authReq struct {
		Password string `json:"password"`
	}
	if err := c.Bind(&authReq); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if authReq.Password != "mytestpass" {
		return c.String(http.StatusUnauthorized, "")
	}
	jwtKey := []byte("secret")
	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString(jwtKey)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	resp := map[string]string{
		"at": signed,
	}
	return c.JSON(http.StatusOK, resp)

}

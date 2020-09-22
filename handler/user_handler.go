package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "Nam",
		"email": "vungocnam@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email string `json:"email"`
		FullName string `json:"full_name"`
		Age int `json:"age"`
	}

	user := User{
		Email: "vungocnam@gmail.com",
		FullName: "Nam",
		Age: 90,
	}
	return c.JSON(http.StatusOK, user)
}
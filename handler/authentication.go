package handler

import (
	"fmt"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Authentication(context echo.Context) error {

	hash, error := bcrypt.GenerateFromPassword([]byte("qwerty"), bcrypt.DefaultCost)

	if error != nil {
		fmt.Println("password hash error")
	}

	hashString := string(hash[:])

	fmt.Println(hashString)

	return context.String(200, "hello....")

}

func (h *Handler) Registration(context echo.Context) error {
	return context.String(200, "registration...")
}

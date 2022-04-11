package main

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID       uint64 `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

var user = User{
	ID:       1,
	UserName: "userName",
	Password: "password",
}

func main() {
	e := echo.New()

	e.POST("login", Login)
	e.Start(":8080")
}

// Login :
func Login(ctx echo.Context) error {
	var u User
	if err := ctx.Bind(&u); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid Json Provided")
		return err
	}

	// compare the user from the request, with the one we defined:
	if user.UserName != u.UserName || user.Password != u.Password {
		ctx.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return nil
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return nil
	}

	return ctx.JSON(http.StatusOK, token)
}

// CreateToken :
func CreateToken(userId uint64) (string, error) {
	var err error

	// Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Member 会員の構造体
type Member struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GetMembers 会員情報取得
func GetMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		members := []Member{
			{Id: 1, Name: "ユーザー1"},
			{Id: 2, Name: "ユーザー2"},
		}
		return c.JSON(http.StatusOK, members)
	}
}

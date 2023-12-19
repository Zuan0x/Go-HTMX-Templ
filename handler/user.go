package handler

import (
	"fmt"

	"github.com/Zuan0x/got/model"
	"github.com/Zuan0x/got/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{Email: "test@test.com"}
	fmt.Println(u)

	return render(c, user.Show(u))
}

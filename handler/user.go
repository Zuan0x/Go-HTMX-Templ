package handler

import (
	"github.com/Zuan0x/got/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	//return render(c, user.Show())
	return user.Show().Render(c.Request().Context(), c.Response())
}

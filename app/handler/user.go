package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/pkg/echoutil"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginRequestBody struct {
	ID       string `json:"id,omitempty" form:"id"`
	Password string `json:"password,omitempty" form:"password"`
}

type Me struct {
	ID string `json:"id,omitempty"`
}

func (h *Handlers) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	newUserReq := LoginRequestBody{}
	if err := c.Bind(&newUserReq); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	fmt.Println(newUserReq)
	if newUserReq.ID == "" || newUserReq.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("項目が空です!"))
	}

	//パスワードをハッシュ
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	newUser := model.User{
		ID:       newUserReq.ID,
		Password: string(hashedPass),
	}

	err = h.Repository.CreateUser(ctx, newUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "success creating a user")
}

func (h *Handlers) Login(c echo.Context) error {
	ctx := c.Request().Context()
	loginReq := LoginRequestBody{}
	if er := c.Bind(&loginReq); er != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, er.Error())
	}
	if loginReq.ID == "" || loginReq.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Item is empty")
	}

	user := model.User{
		ID:       loginReq.ID,
		Password: loginReq.Password,
	}

	user, err := h.Repository.FindUserByID(ctx, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return echo.NewHTTPError(http.StatusForbidden, errors.New("the password is wrong"))
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, err := session.Get(echoutil.SessionStoreKey, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	sess.Values[echoutil.SessionUserIDKey] = loginReq.ID
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, "success loging in")
}

func (h *Handlers) Me(c echo.Context) error {
	userID, err := echoutil.GetUserID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Me{
		ID: userID,
	})
}

// logout
// ログアウトするときにセッションを削除する
func (h *Handlers) Logout(c echo.Context) error {
	sess, err := session.Get(echoutil.SessionStoreKey, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, "success loging out")
}

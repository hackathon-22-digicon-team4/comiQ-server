package echoutil

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SessionStoreKey  = "session"
	SessionUserIDKey = "userID"
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SessionStoreKey, c)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "something wrong in getting session")
		}

		if sess.Values[SessionUserIDKey] == nil {
			return c.String(http.StatusForbidden, "please login")
		}
		c.Set(SessionUserIDKey, sess.Values[SessionUserIDKey].(string))

		return next(c)
	}
}

// session情報からuserIDを取得する
func GetUserID(c echo.Context) (string, error) {
	sess, err := session.Get(SessionStoreKey, c)
	if err != nil {
		return "", err
	}
	return sess.Values[SessionUserIDKey].(string), nil
}

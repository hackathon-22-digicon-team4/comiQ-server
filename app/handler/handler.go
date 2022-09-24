package handler

import (
	"github.com/gorilla/sessions"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/repository"
	"github.com/hackathon-22-digicon-team4/comiQ-server/pkg/echoutil"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	Repository   repository.Repository
	SeesionStore sessions.Store
}

func NewHandlers(repo repository.Repository, sessionStore sessions.Store) *Handlers {
	return &Handlers{
		Repository:   repo,
		SeesionStore: sessionStore,
	}
}

func (h *Handlers) NewServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(h.SeesionStore))
	e.Use(middleware.CORS())

	api := e.Group("/v1")
	apiUsers := api.Group("/users")
	apiUsers.POST("/signup", h.SignUp)
	apiUsers.POST("/login", h.Login)
	apiUsers.GET("/me", h.Me, echoutil.CheckLogin)
	return e
}

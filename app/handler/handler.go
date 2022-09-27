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
	AssetHost    string
}

func NewHandlers(repo repository.Repository, sessionStore sessions.Store, assetHost string) *Handlers {
	return &Handlers{
		Repository:   repo,
		SeesionStore: sessionStore,
		AssetHost:    assetHost,
	}
}

func (h *Handlers) NewServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(h.SeesionStore))
	e.Use(middleware.CORS())

	api := e.Group("/v1")
	api.GET("/", h.Health)
	apiUsers := api.Group("/users")
	apiUsers.POST("/signup", h.SignUp)
	apiUsers.POST("/login", h.Login)
	apiUsers.POST("/logout", h.Logout, echoutil.CheckLogin)
	apiUsers.GET("/me", h.Me, echoutil.CheckLogin)

	apiAuthors := api.Group("/authors")
	apiAuthors.GET("", h.GetAuthors)

	apiBookSeries := api.Group("/book_series")
	apiBookSeries.GET("", h.GetBookSeries)
	apiBookSeries.GET("/:id/books", h.GetBooksByBookSeriesID)

	apiBooks := api.Group("/books")
	apiBooks.GET("/:id", h.GetBookByID)

	apiStamps := api.Group("/stamps")
	apiStamps.GET("", h.GetStamps)

	apiBookUserStamp := api.Group("/book_user_stamps")
	apiBookUserStamp.GET("", h.GetBookUserStamps)
	apiBookUserStamp.POST("", h.PostBookUserStamps, echoutil.CheckLogin)
	apiBookUserStamp.DELETE("/:id", h.DeleteBookUserStamp, echoutil.CheckLogin)
	return e
}

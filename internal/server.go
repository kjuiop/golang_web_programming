package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

const _defaultPort = 8080

type Server struct {
	controller Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	service := NewService(*NewRepository(data))
	controller := NewController(*service)
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()
	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.controller)
}

func RouteMemberships(e *echo.Group, c Controller) {

	e.POST("/memberships", c.Create, middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return "x-My-Request-Header"
		},
	}))

	e.PUT("/memberships", c.Update)

	e.DELETE("/memberships/:id", c.Delete)

	e.GET("/memberships", c.GetAll)

	e.GET("/memberships/:id", c.GetByID)
}

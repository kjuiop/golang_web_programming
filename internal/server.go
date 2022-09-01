package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
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

	g.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Printf(
			"Host : %s , Header : %s , Url : %s , method : %s , status : %d",
			c.Request().Host, c.Request().Header, c.Request().RequestURI, c.Request().Method, c.Response().Status,
		)
		if len(reqBody) > 0 {
			log.Printf("%s \n %s", "request Param : ", string(reqBody))
		}
		if len(resBody) > 0 {
			log.Printf("%s \n %s", "response Body : ", string(resBody))
		}
	}))

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

func middleware1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read origin body bytes
		var bodyBytes []byte
		if c.Request().Body != nil {
			json_map := make(map[string]interface{})

			bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
			// write back to request body
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			err := json.Unmarshal(bodyBytes, &json_map)
			if err != nil {
				return c.JSON(400, "error json.")
			}
			fmt.Println("req : ", json_map)
		}

		return next(c)
	}
}

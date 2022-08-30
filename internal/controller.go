package internal

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrInternalServerError
	}

	res, err := controller.service.Create(req)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	res, err := controller.service.GetByID(id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, res)
}

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(c.Request().Method)
		return next(c)
	}
}

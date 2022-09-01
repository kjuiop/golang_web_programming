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
		return echo.ErrBadRequest
	}

	res, err := controller.service.Create(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) Update(c echo.Context) error {

	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	res, err := controller.service.Update(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) Delete(c echo.Context) error {

	id := c.Param("id")
	if !controller.checkEmptyValue(id) {
		return echo.ErrBadRequest
	}

	err := controller.service.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success delete")
}

func (controller *Controller) GetByID(c echo.Context) error {
	id := c.Param("id")

	if !controller.checkEmptyValue(id) {
		return echo.ErrBadRequest
	}

	res, err := controller.service.GetByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) GetAll(c echo.Context) error {

	res, err := controller.service.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(c.Request().Method)
		return next(c)
	}
}

func (controller *Controller) checkEmptyValue(s string) bool {
	return len(s) > 0
}

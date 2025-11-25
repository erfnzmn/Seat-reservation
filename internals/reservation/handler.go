package reservation

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateReservation)
	g.GET("/:id", h.GetReservation)
	g.PUT("/:id/cancel", h.CancelReservation)
}

func (h *Handler) CreateReservation(c echo.Context) error {
	var input CreateReservationInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid json",
		})
	}

	res, err := h.service.CreateReservation(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}


func (h *Handler) GetReservation(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid id",
		})
	}

	res, err := h.service.GetReservation(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "reservation not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) CancelReservation(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid id",
		})
	}

	if err := h.service.CancelReservation(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "reservation cancelled",
	})
}
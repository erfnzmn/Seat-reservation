package seats

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
	g.GET("", h.GetSeatsByHall)
}

func (h *Handler) GetSeatsByHall(c echo.Context) error {
    hallIDStr := c.Param("hall_id")
    hallID, err := strconv.Atoi(hallIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid hall id"})
    }

    seats, err := h.service.GetSeatsByHall(uint(hallID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, seats)
} 
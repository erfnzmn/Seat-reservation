package waitinglist

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

type joinRequest struct {
	ShowID    uint   `json:"show_id"`
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.Join)
}

func (h *Handler) Join(c echo.Context) error {
	var req joinRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}

	if req.ShowID == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "show_id is required"})
	}

	w, err := h.service.Join(req.ShowID, req.UserName, req.UserPhone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, w)
}

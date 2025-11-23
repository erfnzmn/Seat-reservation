package shows

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service  Service
	adminKey string
}

func NewHandler(service Service, adminKey string) *Handler {
	return &Handler{
		service:  service,
		adminKey: adminKey,
	}
}

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.GetAllShows)
	g.GET("/:id", h.GetShowByID)
	g.POST("", h.CreateShow)
	g.PUT("/:id", h.UpdateShow)
	g.DELETE("/:id", h.DeleteShow)
}

// ---------------------- HANDLERS ----------------------

func (h *Handler) GetAllShows(c echo.Context) error {
	shows, err := h.service.GetAllShows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, shows)
}

func (h *Handler) GetShowByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	show, err := h.service.GetShowByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "show not found"})
	}

	return c.JSON(http.StatusOK, show)
}

func (h *Handler) CreateShow(c echo.Context) error {
	var input CreateShowInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}

	// admin key check
	if c.QueryParam("admin_key") != h.adminKey {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "forbidden"})
	}

	show, err := h.service.CreateShow(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, show)
}

func (h *Handler) UpdateShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var input UpdateShowInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}

	if c.QueryParam("admin_key") != h.adminKey {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "forbidden"})
	}

	updated, err := h.service.UpdateShow(uint(id), input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if c.QueryParam("admin_key") != h.adminKey {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "forbidden"})
	}

	if err := h.service.DeleteShow(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "deleted"})
}

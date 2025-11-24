package halls

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
    g.GET("", h.GetAllHalls)
    g.GET("/:id", h.GetHallByID)
}

func (h *Handler) GetAllHalls(c echo.Context) error {
    halls, err := h.service.GetAllHalls()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, halls)
}

func (h *Handler) GetHallByID(c echo.Context) error {
    idInt, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid hall id"})
    }

    hall, err := h.service.GetHallByID(uint(idInt))
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "hall not found"})
    }

    return c.JSON(http.StatusOK, hall)
}

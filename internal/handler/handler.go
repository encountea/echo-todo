package handler

import (
	"github.com/encountea/echo-todo/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *echo.Echo {
	r := echo.New()

	r.GET("/", h.getAllTasks)
	r.POST("/", h.createTask)
	r.GET("/:id", h.getTaskById)
	r.PUT("/:id", h.updateTask)
	r.DELETE("/:id", h.deleteTask)

	return r
}

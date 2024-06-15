package handler

import (
	"net/http"

	models "github.com/encountea/echo-todo/internal"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getAllTasks(c echo.Context) error {
	tasks, err := h.services.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getTaskById(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	task, err := h.services.GetByUUID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (h *Handler) createTask(c echo.Context) error {
	var task models.Task
    if err := c.Bind(&task); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    if err := h.services.Create(task); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, task)
}

func (h *Handler) updateTask(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid UUID")
    }

    var task models.Task
    if err := c.Bind(&task); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid input")
    }

    task.UUID = id
    if err := h.services.Update(task); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, task)
}

func (h *Handler) deleteTask(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
   if err != nil {
       return c.JSON(http.StatusBadRequest, "Invalid UUID")
   }

    if err := h.services.Delete(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.NoContent(http.StatusNoContent)
}

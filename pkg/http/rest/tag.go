package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllTags handler returns a list of all tags
func (h *Handler) GetAllTags(c echo.Context) error {
	tags, err := h.lister.AllTags()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tags)
}

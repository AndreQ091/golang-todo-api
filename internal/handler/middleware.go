package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)

}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get("userId")

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return 0, errors.New("user not found")
	}

	idInt, ok := userId.(int)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid user id type")
		return 0, errors.New("invalid user id type")
	} 

	return idInt, nil
}

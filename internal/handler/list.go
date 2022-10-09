package handler

import (
	"net/http"
	"strconv"

	todo "github.com/AndreQ091/golang-todo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateList(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	var input todo.TodoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetLists(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": lists,
	})
}

func (h *Handler) GetOneList(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	list, err := h.services.TodoList.GetById(id, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}

func (h *Handler) UpdateList(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	var input todo.UpdateListInput

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.UpdateById(id, userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": nil,
	})
}

func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.TodoList.DeleteById(id, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": nil,
	})
}

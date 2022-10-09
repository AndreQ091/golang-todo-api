package handler

import (
	"net/http"
	"strconv"

	todo "github.com/AndreQ091/golang-todo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateItem(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id")
		return
	}

	var input todo.TodoItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(listId, userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetItems(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id")
		return
	}

	lists, err := h.services.TodoItem.GetAll(listId, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": lists,
	})
}

func (h *Handler) GetOneItem(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	list, err := h.services.TodoItem.GetById(id, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}

func (h *Handler) UpdateItem(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	var input todo.UpdateItemInput

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.UpdateById(id, userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": nil,
	})
}

func (h *Handler) DeleteItem(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.TodoItem.DeleteById(id, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": nil,
	})
}

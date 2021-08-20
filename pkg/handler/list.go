package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/VladMak/ServiceAPI"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	fmt.Println(id)

	var input ServiceAPI.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// call service method
}

func (h *Handler) getAllLists(c *gin.Context) {
	
}

func (h *Handler) getListById(c *gin.Context) {
	
}

func (h *Handler) updateList(c *gin.Context) {
	
}

func (h *Handler) deleteList(c *gin.Context) {
	
}
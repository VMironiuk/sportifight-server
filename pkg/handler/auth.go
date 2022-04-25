package handler

import (
	"net/http"

	"github.com/VMironiuk/sportifight-server"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input sportifight.User
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	// TODO: add implementation
	ctx.JSON(http.StatusOK, map[string]string{
		"ok": "sign-in test response",
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	// TODO: add implementation
	ctx.JSON(http.StatusOK, map[string]string{
		"ok": "sign-up test response",
	})
}

func (h *Handler) signIn(ctx *gin.Context) {
	// TODO: add implementation
	ctx.JSON(http.StatusOK, map[string]string{
		"ok": "sign-in test response",
	})
}

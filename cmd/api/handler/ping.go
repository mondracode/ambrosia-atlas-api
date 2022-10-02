package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct {
}

func (handler Ping) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

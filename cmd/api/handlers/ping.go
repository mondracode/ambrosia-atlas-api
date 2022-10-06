package handlers

import (
	"github.com/gin-gonic/gin"
)

type Ping struct {
}

func NewPing() *Ping {
	return &Ping{}
}

func (handler Ping) Ping(c *gin.Context) (interface{}, error) {
	return "pong", nil
}

package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Handler(ginCtx *gin.Context)
}

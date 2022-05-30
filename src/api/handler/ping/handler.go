package ping

import (
	"EurekaHome/src/api/handler/ping/contract"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	pong := contract.Pong{
		Message: "Eureka Pong",
	}
	ginCTX.JSON(http.StatusOK, pong)
}


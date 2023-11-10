package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) ListVirtualAccount(c *gin.Context) {
	virtualAccounts, err := server.q.ListVirtualAccount(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    virtualAccounts,
	})
}

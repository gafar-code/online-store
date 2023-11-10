package api

import (
	"net/http"
	"time"

	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gin-gonic/gin"
)

type listTransactionRes struct {
	Page  int32                 `json:"page"`
	Size  int32                 `json:"size"`
	Items []listTransactionItem `json:"items"`
}

type listTransactionItem struct {
	ID        int64          `json:"id"`
	Status    string         `json:"status"`
	Amount    int64          `json:"amount"`
	IssuedAt  time.Time      `json:"issued_at"`
	CreatedAt time.Time      `json:"created_at"`
	Products  []db.OrderItem `json:"products"`
}

func (server *Server) ListTransaction(c *gin.Context, params ListTransactionParams) {
	var req PaginationReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	customer, err := getCustomerByToken(server, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	arg := db.ListTransactionByCustomerIdParams{
		CustomerID: customer.ID,
		Limit:      req.Size,
		Offset:     (req.Page - 1) * req.Size,
	}

	transactions, err := server.q.ListTransactionByCustomerId(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	items := []listTransactionItem{}

	for _, trx := range transactions {
		order, err := server.q.GetOrder(c, trx.OrderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		orderItems, err := server.q.ListOrderItemByOrderId(c, order.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		items = append(items, listTransactionItem{
			ID:        trx.ID,
			Status:    trx.Status,
			Amount:    order.Amount,
			IssuedAt:  trx.IssuedAt,
			CreatedAt: trx.CreatedAt,
			Products:  orderItems,
		})

	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: "Success",
		Data: listTransactionRes{
			Page:  req.Page,
			Size:  req.Size,
			Items: items,
		},
	})

}

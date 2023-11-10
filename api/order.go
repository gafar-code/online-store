package api

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gin-gonic/gin"
)

const paymentExpired = 24 * time.Hour

type addOrderReq struct {
	VirtualAccountID int64          `json:"virtual_account_id" binding:"required"`
	Amount           int64          `json:"amount" binding:"required"`
	OrderItems       []addOrderItem `json:"products" binding:"required"`
}

type addOrderItem struct {
	ProductID    int64 `json:"product_id" binding:"required"`
	ProductPrice int64 `json:"product_price" binding:"required"`
	Qty          int64 `json:"qty" binding:"required"`
}

type addOrderRes struct {
	ID             int64     `json:"id"`
	BankName       string    `json:"bank_name"`
	RekeningNumber int64     `json:"rekening_number"`
	Amount         int64     `json:"amount"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	IssuedAt       time.Time `json:"issued_at"`
	ExpiredAt      time.Time `json:"expired_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type orderProofReq struct {
	OrderID        int64                 `form:"order_id"`
	Name           string                `form:"name"`
	RekeningNumber int64                 `form:"rekening_number"`
	Image          *multipart.FileHeader `form:"image"`
}

type updateOrderProofReq struct {
	OrderID int64 `form:"order_id"`
}

type orderProofRes struct {
	ID             int64     `json:"id"`
	BankName       string    `json:"bank_name"`
	RekeningNumber int64     `json:"rekening_number"`
	Amount         int64     `json:"amount"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	IssuedAt       time.Time `json:"issued_at"`
	ExpiredAt      time.Time `json:"expired_at"`
	CreatedAt      time.Time `json:"created_at"`
}

func (server *Server) AddOrder(c *gin.Context) {
	var req addOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
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

	arg := db.CreateOrderParams{
		CustomerID:       customer.ID,
		Amount:           req.Amount,
		Status:           "WAITING_PAYMENT",
		VirtualAccountID: req.VirtualAccountID,
		ExpiredAt:        time.Now().Add(paymentExpired),
	}

	order, err := server.q.CreateOrder(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	for _, item := range req.OrderItems {
		prod, err := server.q.GetProductDetail(c, item.ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		arg := db.CreateOrderItemParams{
			CategoryID:  prod.CategoryID,
			Name:        prod.Name,
			ImageUrl:    prod.ImageUrl,
			Description: prod.Description,
			Price:       item.ProductPrice,
			Qty:         item.Qty,
			ProductID:   item.ProductID,
			OrderID:     order.ID,
		}

		_, err = server.q.CreateOrderItem(c, arg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}

	va, err := server.q.GetVirtualAccount(c, req.VirtualAccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	result := addOrderRes{
		ID:             order.ID,
		BankName:       va.Name,
		RekeningNumber: va.RekeningNumber,
		Amount:         req.Amount,
		Description:    va.Description,
		Status:         order.Status,
		IssuedAt:       order.IssuedAt,
		ExpiredAt:      order.ExpiredAt,
		CreatedAt:      order.CreatedAt,
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (server *Server) AddProofPayment(c *gin.Context) {
	var req orderProofReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	orderId, err := strconv.ParseInt(c.PostForm("order_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	rekeningNumb, err := strconv.ParseInt(c.PostForm("rekening_number"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req = orderProofReq{
		OrderID:        orderId,
		Name:           c.PostForm(req.Name),
		RekeningNumber: rekeningNumb,
		Image:          req.Image,
	}

	customer, err := getCustomerByToken(server, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	fileName := fmt.Sprintf("%v_%d_%s", customer.Name, time.Now().UnixNano(), file.Filename)
	filePath := "images/proofs/" + fileName

	fmt.Println(c.PostForm("order_id"))

	arg := db.CreateOrderProofParams{
		OrderID:        req.OrderID,
		NameHolder:     req.Name,
		RekeningNumber: req.RekeningNumber,
		ImageUrl:       filePath,
	}

	proof, err := server.q.CreateOrderProof(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	updateArg := db.UpdateOrderParams{
		Status: "PENDING",
		ID:     proof.OrderID,
	}

	updateOrder, err := server.q.UpdateOrder(c, updateArg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	order, err := server.q.GetOrder(c, proof.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	va, err := server.q.GetVirtualAccount(c, order.VirtualAccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	result := orderProofRes{
		ID:             order.ID,
		BankName:       va.Name,
		RekeningNumber: va.RekeningNumber,
		Amount:         updateOrder.Amount,
		Description:    va.Description,
		Status:         updateOrder.Status,
		IssuedAt:       updateOrder.IssuedAt,
		ExpiredAt:      updateOrder.ExpiredAt,
		CreatedAt:      updateOrder.CreatedAt,
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (server *Server) UpdateOrderProof(c *gin.Context, params UpdateOrderProofParams) {
	var req updateOrderProofReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	arg := db.UpdateOrderParams{
		ID:     req.OrderID,
		Status: "PAID",
	}

	order, err := server.q.UpdateOrder(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	trxArg := db.CreateTransactionParams{
		CustomerID: order.CustomerID,
		Status:     "ON_PROGRESS",
		IssuedAt:   order.IssuedAt,
		OrderID:    req.OrderID,
	}

	_, err = server.q.CreateTransaction(c, trxArg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success",
	})

}

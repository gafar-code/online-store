package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gafar-code/online-store/util"
	"github.com/gin-gonic/gin"
)

type addCartReq struct {
	ProductID int64 `json:"product_id" binding:"required"`
}

type deleteCartReq struct {
	ProductID int64 `form:"product_id" binding:"required"`
}

type bulkDeleteCartReq struct {
	ProductIDs []int64 `form:"product_ids" binding:"required"`
}

type listCartRes struct {
	Page  int32         `json:"page"`
	Size  int32         `json:"size"`
	Items []cartItemRes `json:"items"`
}

type cartItemRes struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Qty         int64     `json:"qty"`
	CreatedAt   time.Time `json:"created_at"`
}

func (server *Server) AddToCart(c *gin.Context) {
	var req addCartReq
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

	arg := db.GetExistingCartParams{
		CustomerID: customer.ID,
		ProductID:  req.ProductID,
	}

	existing, err := server.q.GetExistingCart(c, arg)
	if err != nil && err == sql.ErrNoRows {
		createArg := db.CreateCartParams{
			CustomerID: customer.ID,
			ProductID:  req.ProductID,
			Qty:        1,
		}

		cart, err := server.q.CreateCart(c, createArg)

		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, ResponseData{
			Code:    http.StatusCreated,
			Message: "Success",
			Data:    cart,
		})

		return
	}

	if !util.IsEmpty(existing) {
		arg := db.UpdateCartParams{
			ID:  existing.ID,
			Qty: existing.Qty + 1,
		}

		cart, err := server.q.UpdateCart(c, arg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, ResponseData{
			Code:    http.StatusCreated,
			Message: "Success",
			Data:    cart,
		})
		return
	}

}

func (server *Server) DeleteProductFromCart(c *gin.Context, params DeleteProductFromCartParams) {
	var req deleteCartReq
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

	arg := db.GetExistingCartParams{
		CustomerID: customer.ID,
		ProductID:  req.ProductID,
	}

	cart, err := server.q.GetExistingCart(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if cart.Qty == 1 {
		err = server.q.DeleteCart(c, cart.ID)
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
		return
	}

	updateArg := db.UpdateCartParams{
		ID:  cart.ID,
		Qty: cart.Qty - 1,
	}

	cart, err = server.q.UpdateCart(c, updateArg)
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
		Data:    cart,
	})

}

func (server *Server) BulkDeleteCart(c *gin.Context, params BulkDeleteCartParams) {
	var req bulkDeleteCartReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for _, id := range req.ProductIDs {
		err := server.q.DeleteCart(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success",
	})
}

func (server *Server) GetCart(c *gin.Context, params GetCartParams) {
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

	arg := db.GetCartByCustomerIdParams{
		CustomerID: customer.ID,
		Limit:      req.Size,
		Offset:     (req.Page - 1) * req.Size,
	}

	carts, err := server.q.GetCartByCustomerId(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	products := []cartItemRes{}

	for _, cart := range carts {
		product, err := server.q.GetProductDetail(c, cart.ProductID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		cartItem := cartItemRes{
			ID:          product.ID,
			CategoryID:  product.CategoryID,
			Name:        product.Name,
			ImageUrl:    product.ImageUrl,
			Description: product.Description,
			Price:       product.Price,
			Qty:         cart.Qty,
			CreatedAt:   product.CreatedAt,
		}

		products = append(products, cartItem)
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: "Success",
		Data: listCartRes{
			Size:  req.Size,
			Page:  req.Page,
			Items: products,
		},
	})
}

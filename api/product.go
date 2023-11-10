package api

import (
	"net/http"

	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gafar-code/online-store/util"
	"github.com/gin-gonic/gin"
)

type listProductReq struct {
	Size       int32 `form:"size" binding:"required"`
	Page       int32 `form:"page" binding:"required"`
	CategoryId int32 `form:"category_id" binding:"omitempty"`
}

type productDetailReq struct {
	ID int64 `uri:"id" binding:"required"`
}

type productListRes struct {
	Page  int32        `json:"page"`
	Size  int32        `json:"size"`
	Items []db.Product `json:"items"`
}

func (server *Server) ListProduct(c *gin.Context, params ListProductParams) {
	var req listProductReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if util.IsEmpty(req.CategoryId) {
		arg := db.ListProductParams{
			Limit:  req.Size,
			Offset: (req.Page - 1) * req.Size,
		}

		products, err := server.q.ListProduct(c, arg)

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
			Data: productListRes{
				Size:  req.Size,
				Page:  req.Page,
				Items: products,
			},
		})

		return
	}

	arg := db.GetProductByCategoryIdParams{
		Limit:      req.Size,
		Offset:     (req.Page - 1) * req.Size,
		CategoryID: int64(req.CategoryId),
	}

	products, err := server.q.GetProductByCategoryId(c, arg)

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
		Data: productListRes{
			Size:  req.Size,
			Page:  req.Page,
			Items: products,
		},
	})
}
func (server *Server) GetProductDetail(c *gin.Context, id int) {
	var req productDetailReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	product, err := server.q.GetProductDetail(c, req.ID)

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
		Data:    product,
	})
}

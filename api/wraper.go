package api

import (
	"github.com/gin-gonic/gin"
)

type ServerWraper struct {
	server *Server
}

// (POST /auth/login)
func (wraper *ServerWraper) Login(c *gin.Context) {
	wraper.server.Login(c)
}

// (POST /auth/register)
func (wraper *ServerWraper) Register(c *gin.Context) {
	wraper.server.Register(c)
}

// (DELETE /cart)
func (wraper *ServerWraper) DeleteProductFromCart(c *gin.Context, params DeleteProductFromCartParams) {

}

// (GET /cart)
func (wraper *ServerWraper) GetCart(c *gin.Context, params GetCartParams) {}

// (POST /cart)
func (wraper *ServerWraper) AddToCart(c *gin.Context) {}

// (DELETE /cart/delete-all)
func (wraper *ServerWraper) DeleteAllCart(c *gin.Context) {}

// (POST /order)
func (wraper *ServerWraper) AddOrder(c *gin.Context) {

}

// (PUT /order/{id})
func (wraper *ServerWraper) UpdatePayment(c *gin.Context, id int) {}

// (GET /product)
func (wraper *ServerWraper) ListProduct(c *gin.Context, params ListProductParams) {
	wraper.server.ListProduct(c)
}

// (GET /product/{id})
func (wraper *ServerWraper) GetProductDetail(c *gin.Context, id int) {
	wraper.server.GetProductDetail(c)
}

// (GET /transaction)
func (wraper *ServerWraper) ListOrders(c *gin.Context, params ListOrdersParams) {}

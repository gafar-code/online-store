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
	wraper.server.DeleteProductFromCart(c)
}

// (GET /cart)
func (wraper *ServerWraper) GetCart(c *gin.Context, params GetCartParams) {
	wraper.server.GetCart(c)
}

// (POST /cart)
func (wraper *ServerWraper) AddToCart(c *gin.Context) {
	wraper.server.AddToCart(c)
}

// (DELETE /cart/bulk)
func (wraper *ServerWraper) BulkDeleteCart(c *gin.Context, params BulkDeleteCartParams) {
	wraper.server.BulkDeleteCart(c)
}

// (POST /order)
func (wraper *ServerWraper) AddOrder(c *gin.Context) {
	wraper.server.AddOrder(c)
}

// (PUT /order/{id})
func (wraper *ServerWraper) AddProofPayment(c *gin.Context) {
	wraper.server.AddProofPayment(c)
}

// (PUT /order/proof)
func (wraper *ServerWraper) UpdateOrderProof(c *gin.Context, params UpdateOrderProofParams) {
	wraper.server.UpdateOrderProof(c)
}

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

// (GET /virtual-account)
func (wraper *ServerWraper) ListVirtualAccount(c *gin.Context) {
	wraper.server.ListVirtualAccount(c)
}

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /auth/login)
	Login(c *gin.Context)

	// (POST /auth/register)
	Register(c *gin.Context)

	// (DELETE /cart)
	DeleteProductFromCart(c *gin.Context, params DeleteProductFromCartParams)

	// (GET /cart)
	GetCart(c *gin.Context, params GetCartParams)

	// (POST /cart)
	AddToCart(c *gin.Context)

	// (DELETE /cart/bulk)
	BulkDeleteCart(c *gin.Context, params BulkDeleteCartParams)

	// (POST /order)
	AddOrder(c *gin.Context)

	// (POST /order/proof)
	AddProofPayment(c *gin.Context)

	// (PUT /order/proof)
	UpdateOrderProof(c *gin.Context, params UpdateOrderProofParams)

	// (GET /product)
	ListProduct(c *gin.Context, params ListProductParams)

	// (GET /product/{id})
	GetProductDetail(c *gin.Context, id int)

	// (GET /transaction)
	ListOrders(c *gin.Context, params ListOrdersParams)

	// (GET /virtual-account)
	ListVirtualAccount(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// Login operation middleware
func (siw *ServerInterfaceWrapper) Login(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Login(c)
}

// Register operation middleware
func (siw *ServerInterfaceWrapper) Register(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Register(c)
}

// DeleteProductFromCart operation middleware
func (siw *ServerInterfaceWrapper) DeleteProductFromCart(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteProductFromCartParams

	// ------------- Required query parameter "product_id" -------------

	if paramValue := c.Query("product_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument product_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "product_id", c.Request.URL.Query(), &params.ProductId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteProductFromCart(c, params)
}

// GetCart operation middleware
func (siw *ServerInterfaceWrapper) GetCart(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCartParams

	// ------------- Required query parameter "page" -------------

	if paramValue := c.Query("page"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "size" -------------

	if paramValue := c.Query("size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "size", c.Request.URL.Query(), &params.Size)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter size: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCart(c, params)
}

// AddToCart operation middleware
func (siw *ServerInterfaceWrapper) AddToCart(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AddToCart(c)
}

// BulkDeleteCart operation middleware
func (siw *ServerInterfaceWrapper) BulkDeleteCart(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params BulkDeleteCartParams

	// ------------- Required query parameter "product_ids" -------------

	if paramValue := c.Query("product_ids"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument product_ids is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "product_ids", c.Request.URL.Query(), &params.ProductIds)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter product_ids: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.BulkDeleteCart(c, params)
}

// AddOrder operation middleware
func (siw *ServerInterfaceWrapper) AddOrder(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AddOrder(c)
}

// AddProofPayment operation middleware
func (siw *ServerInterfaceWrapper) AddProofPayment(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AddProofPayment(c)
}

// UpdateOrderProof operation middleware
func (siw *ServerInterfaceWrapper) UpdateOrderProof(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdateOrderProofParams

	// ------------- Required query parameter "order_id" -------------

	if paramValue := c.Query("order_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument order_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "order_id", c.Request.URL.Query(), &params.OrderId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter order_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateOrderProof(c, params)
}

// ListProduct operation middleware
func (siw *ServerInterfaceWrapper) ListProduct(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListProductParams

	// ------------- Required query parameter "page" -------------

	if paramValue := c.Query("page"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "size" -------------

	if paramValue := c.Query("size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "size", c.Request.URL.Query(), &params.Size)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter size: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "category_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "category_id", c.Request.URL.Query(), &params.CategoryId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter category_id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListProduct(c, params)
}

// GetProductDetail operation middleware
func (siw *ServerInterfaceWrapper) GetProductDetail(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProductDetail(c, id)
}

// ListOrders operation middleware
func (siw *ServerInterfaceWrapper) ListOrders(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListOrdersParams

	// ------------- Required query parameter "page" -------------

	if paramValue := c.Query("page"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument page is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "size" -------------

	if paramValue := c.Query("size"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument size is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "size", c.Request.URL.Query(), &params.Size)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter size: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListOrders(c, params)
}

// ListVirtualAccount operation middleware
func (siw *ServerInterfaceWrapper) ListVirtualAccount(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ListVirtualAccount(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/auth/login", wrapper.Login)
	router.POST(options.BaseURL+"/auth/register", wrapper.Register)
	router.DELETE(options.BaseURL+"/cart", wrapper.DeleteProductFromCart)
	router.GET(options.BaseURL+"/cart", wrapper.GetCart)
	router.POST(options.BaseURL+"/cart", wrapper.AddToCart)
	router.DELETE(options.BaseURL+"/cart/bulk", wrapper.BulkDeleteCart)
	router.POST(options.BaseURL+"/order", wrapper.AddOrder)
	router.POST(options.BaseURL+"/order/proof", wrapper.AddProofPayment)
	router.PUT(options.BaseURL+"/order/proof", wrapper.UpdateOrderProof)
	router.GET(options.BaseURL+"/product", wrapper.ListProduct)
	router.GET(options.BaseURL+"/product/:id", wrapper.GetProductDetail)
	router.GET(options.BaseURL+"/transaction", wrapper.ListOrders)
	router.GET(options.BaseURL+"/virtual-account", wrapper.ListVirtualAccount)
}

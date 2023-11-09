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

	// (DELETE /cart/delete-all)
	DeleteAllCart(c *gin.Context)

	// (POST /order)
	AddOrder(c *gin.Context)

	// (PUT /order/{id})
	UpdatePayment(c *gin.Context, id int)

	// (GET /product)
	GetProductByCategoryId(c *gin.Context, params GetProductByCategoryIdParams)

	// (GET /product/{id})
	GetProductDetail(c *gin.Context, id int)

	// (GET /transaction)
	ListOrders(c *gin.Context, params ListOrdersParams)
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

	// ------------- Required query parameter "cart_id" -------------

	if paramValue := c.Query("cart_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument cart_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "cart_id", c.Request.URL.Query(), &params.CartId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter cart_id: %w", err), http.StatusBadRequest)
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

// DeleteAllCart operation middleware
func (siw *ServerInterfaceWrapper) DeleteAllCart(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteAllCart(c)
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

// UpdatePayment operation middleware
func (siw *ServerInterfaceWrapper) UpdatePayment(c *gin.Context) {

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

	siw.Handler.UpdatePayment(c, id)
}

// GetProductByCategoryId operation middleware
func (siw *ServerInterfaceWrapper) GetProductByCategoryId(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProductByCategoryIdParams

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

	siw.Handler.GetProductByCategoryId(c, params)
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
	router.DELETE(options.BaseURL+"/cart/delete-all", wrapper.DeleteAllCart)
	router.POST(options.BaseURL+"/order", wrapper.AddOrder)
	router.PUT(options.BaseURL+"/order/:id", wrapper.UpdatePayment)
	router.GET(options.BaseURL+"/product", wrapper.GetProductByCategoryId)
	router.GET(options.BaseURL+"/product/:id", wrapper.GetProductDetail)
	router.GET(options.BaseURL+"/transaction", wrapper.ListOrders)
}

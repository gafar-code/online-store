// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Customer defines model for Customer.
type Customer struct {
	Address   *string              `json:"address,omitempty"`
	CreatedAt *time.Time           `json:"created_at,omitempty"`
	Email     *openapi_types.Email `json:"email,omitempty"`
	Id        *int32               `json:"id,omitempty"`
	Name      *string              `json:"name,omitempty"`
	Token     *string              `json:"token,omitempty"`
}

// Order defines model for Order.
type Order struct {
	Amount         *int       `json:"amount,omitempty"`
	BankName       *string    `json:"bank_name,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	Description    *string    `json:"description,omitempty"`
	ExpiredAt      *time.Time `json:"expired_at,omitempty"`
	Id             *int       `json:"id,omitempty"`
	IssuedAt       *time.Time `json:"issued_at,omitempty"`
	RekeningNumber *int       `json:"rekening_number,omitempty"`
	Status         *string    `json:"status,omitempty"`
}

// Product defines model for Product.
type Product struct {
	CategoryId  *int32     `json:"category_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *int32     `json:"id,omitempty"`
	ImageUrl    *string    `json:"image_url,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Price       *int       `json:"price,omitempty"`
	Qty         *int       `json:"qty,omitempty"`
}

// Transaction defines model for Transaction.
type Transaction struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        *string    `json:"id,omitempty"`
	IssuedAt  *time.Time `json:"issued_at,omitempty"`
	Products  *[]Product `json:"products,omitempty"`
	Status    *string    `json:"status,omitempty"`
}

// LoginJSONBody defines parameters for Login.
type LoginJSONBody struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// RegisterJSONBody defines parameters for Register.
type RegisterJSONBody struct {
	Address  *string             `json:"address,omitempty"`
	Email    openapi_types.Email `json:"email"`
	Name     *string             `json:"name,omitempty"`
	Password string              `json:"password"`
}

// DeleteProductFromCartParams defines parameters for DeleteProductFromCart.
type DeleteProductFromCartParams struct {
	CartId int `form:"cart_id" json:"cart_id"`
}

// GetCartParams defines parameters for GetCart.
type GetCartParams struct {
	// Page The page number you want to retrieve, starting from page 1
	Page int `form:"page" json:"page"`

	// Size Number of items per page
	Size int `form:"size" json:"size"`
}

// AddToCartJSONBody defines parameters for AddToCart.
type AddToCartJSONBody struct {
	ProductId *int `json:"product_id,omitempty"`
	Qty       *int `json:"qty,omitempty"`
}

// AddOrderJSONBody defines parameters for AddOrder.
type AddOrderJSONBody struct {
	Products *[]struct {
		ProductId *int `json:"product_id,omitempty"`
		Qty       *int `json:"qty,omitempty"`
	} `json:"products,omitempty"`
	VirtualAccountId int `json:"virtual_account_id"`
}

// UpdatePaymentJSONBody defines parameters for UpdatePayment.
type UpdatePaymentJSONBody struct {
	ImageUrl         *string `json:"image_url,omitempty"`
	Name             *string `json:"name,omitempty"`
	RekeningNumber   *int    `json:"rekening_number,omitempty"`
	VirtualAccountId int     `json:"virtual_account_id"`
}

// ListProductParams defines parameters for ListProduct.
type ListProductParams struct {
	// Page The page number you want to retrieve, starting from page 1
	Page int `form:"page" json:"page"`

	// Size Number of items per page
	Size int `form:"size" json:"size"`

	// CategoryId View product list by product category
	CategoryId *int `form:"category_id,omitempty" json:"category_id,omitempty"`
}

// ListOrdersParams defines parameters for ListOrders.
type ListOrdersParams struct {
	// Page The page number you want to retrieve, starting from page 1
	Page int `form:"page" json:"page"`

	// Size Number of items per page
	Size int `form:"size" json:"size"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody RegisterJSONBody

// AddToCartJSONRequestBody defines body for AddToCart for application/json ContentType.
type AddToCartJSONRequestBody AddToCartJSONBody

// AddOrderJSONRequestBody defines body for AddOrder for application/json ContentType.
type AddOrderJSONRequestBody AddOrderJSONBody

// UpdatePaymentJSONRequestBody defines body for UpdatePayment for application/json ContentType.
type UpdatePaymentJSONRequestBody UpdatePaymentJSONBody

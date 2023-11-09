package api

import (
	"net/http"
	"time"

	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gafar-code/online-store/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type registerReq struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type customerResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

const tokenDuration = 7 * (24 * time.Hour)

func (server *Server) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	customer, err := server.q.GetCustomerByEmail(c, req.Email)

	if err != nil {
		if util.IsEmpty(customer) {
			c.JSON(http.StatusNotFound, Response{
				Code:    http.StatusNotFound,
				Message: "Email Tidak Terdaftar",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "Password Salah",
		})
		return
	}

	token, err := tokenMaker.CreateToken(customer.Email, tokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: "Login Berhasil!",
		Data: customerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			Email:     customer.Email,
			Token:     token,
			CreatedAt: customer.CreatedAt,
		},
	})
}

func (server *Server) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	customer, _ := server.q.GetCustomerByEmail(c, req.Email)
	if !util.IsEmpty(customer) {
		c.JSON(http.StatusConflict, Response{
			Code:    http.StatusConflict,
			Message: "Email Sudah Terdaftar",
		})
		return

	}

	token, err := tokenMaker.CreateToken(customer.Email, tokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	arg := db.CreateCustomerParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Address:  req.Address,
	}

	customer, err = server.q.CreateCustomer(c, arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ResponseData{
		Code:    http.StatusCreated,
		Message: "Registrasi Berhasil!",
		Data: customerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			Email:     customer.Email,
			Address:   customer.Address,
			Token:     token,
			CreatedAt: customer.CreatedAt,
		},
	})
}

package api

import (
	"database/sql"
	"errors"
	"fmt"

	db "github.com/gafar-code/online-store/db/sqlc"
	"github.com/gafar-code/online-store/token"
	"github.com/gafar-code/online-store/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	q          *db.Queries
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

type ResponseData struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type PaginationReq struct {
	Page int32 `form:"page" binding:"required,min=1"`
	Size int32 `form:"size" binding:"required"`
}

var tokenMaker token.Maker

func errorHandler(c *gin.Context, err error, code int) {
	if err != nil {
		c.JSON(code, Response{
			Code:    int32(code),
			Message: err.Error(),
		})
		return
	}
}

func getCustomerByToken(server *Server, c *gin.Context) (customer db.Customer, err error) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	err = authPayload.Valid()
	if err != nil {
		return
	}

	customer, err = server.q.GetCustomerByEmail(c, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("invalid token")
			return
		}
		return
	}

	return
}

func NewServer(config util.Config, q *db.Queries) (server *Server, err error) {
	tokenMaker, err = token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	router := gin.Default()
	server = &Server{
		q:          q,
		router:     router,
		tokenMaker: tokenMaker,
		config:     config,
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "api_key", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	middleware := []MiddlewareFunc{
		AuthMiddleware,
	}

	baseUrl := fmt.Sprintf("/api/%v", config.APIVersion)

	RegisterHandlersWithOptions(router, server, GinServerOptions{
		BaseURL:      baseUrl,
		Middlewares:  middleware,
		ErrorHandler: errorHandler,
	})

	server.router = router

	return
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

package api

import (
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

var tokenMaker token.Maker

func errorHandler(c *gin.Context, err error, code int) {
	if err != nil {
		c.JSON(code, Response{
			Code:    int32(code),
			Message: err.Error(),
		})
	}
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

	baseUrl := "/api/v1"
	wrapServer := &ServerWraper{
		server: server,
	}

	// TODO: Change wrapServer to server
	RegisterHandlersWithOptions(router, wrapServer, GinServerOptions{
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

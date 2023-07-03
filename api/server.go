package api

import (
	db "Bank/db/sqlc"
	"Bank/token"
	"Bank/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", s.createUser)
	router.GET("/users/:username", s.getUser)
	router.POST("/users/login", s.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))

	authRoutes.POST("/accounts", s.createAccount)
	authRoutes.GET("/accounts/:id", s.getAccount)
	authRoutes.GET("/accounts", s.listAccounts)
	authRoutes.POST("/accounts/update", s.updateAccount)
	authRoutes.GET("/accounts/delete/:id", s.deleteAccount)

	authRoutes.POST("/transfers", s.createTransfer)

	s.router = router
}

// Start starts the HTTP server on given address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

package apihandler

import (
	"fmt"

	"github.com/KhanSufiyanMirza/mini-aspire-API/db"
	"github.com/KhanSufiyanMirza/mini-aspire-API/token"
	"github.com/KhanSufiyanMirza/mini-aspire-API/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     utils.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store *db.Store) (*Server, error) {
	tokMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("connot create token maker : %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//for login
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	//here we check person is authenticated or not
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/:id", server.getUser)
	authRoutes.GET("/users", server.listUsers)

	authRoutes.POST("/loan/createLoan", server.createLoan)
	authRoutes.GET("/loan/:id", server.getLoan)
	authRoutes.GET("/loan", server.listLoan)

	authRoutes.POST("/payment", server.createPayment)
	authRoutes.GET("/payment/:id", server.getPayment)
	authRoutes.GET("/payment", server.listPayment)

	// router.GET("/requests/:id", server.getRequest)
	// router.GET("/requests", server.listRequest)
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/shama3541/simplebank/db/database"
	"github.com/shama3541/simplebank/token"
	"github.com/shama3541/simplebank/util"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store *db.Store) *Server {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	server.router = gin.Default()
	server.router.POST("/user", server.CreateUserHandler)
	server.router.POST("/login", server.UserLoginHanderl)
	authRoutes := server.router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/accounts", server.GetAllAccounts)
	authRoutes.GET("/accounts/:id", server.GetAccountByID)
	authRoutes.POST("/transfer", server.CreateTransfer)

	return server

}

func (server *Server) StartServer(address string) error {
	return server.router.Run()
}

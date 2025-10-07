package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/shama3541/simplebank/db/database"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	server.router = gin.Default()

	server.router.POST("/accounts", server.CreateAccount)
	server.router.GET("/accounts", server.GetAllAccounts)
	server.router.GET("/accounts/:id", server.GetAccountByID)
	server.router.POST("/transfer", server.CreateTransfer)
	server.router.POST("/user", server.CreateUserHandler)
	server.router.POST("/login", server.UserLoginHanderl)
	return server

}

func (server *Server) StartServer(address string) error {
	return server.router.Run()
}

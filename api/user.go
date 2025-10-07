package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/shama3541/simplebank/db/database"
	"github.com/shama3541/simplebank/util"
)

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (server *Server) CreateUserHandler(ctx *gin.Context) {

	var req UserRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "Bad request or invalid params",
			"error":   err.Error(),
		})
		return

	}
	hashedpassword, _ := util.HashedPassword(req.Password)

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedpassword,
		Email:          req.Email,
		FullName:       req.FullName,
	}

	resp, err := server.store.CreateUser(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, resp)

}

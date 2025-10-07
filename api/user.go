package api

import (
	"database/sql"
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

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) UserLoginHanderl(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Badrequest",
			"error":   err.Error(),
		})
		return
	}

	resp, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found please register the account",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return

	}

	err = util.CheckHashesPassword(resp.HashedPassword, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Wrong password",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User is able to login ",
	})

}

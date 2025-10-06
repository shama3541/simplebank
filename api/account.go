package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/shama3541/simplebank/db/database"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to fetch user",
			"details": err.Error(),
		})
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	accountresp, err := server.store.Queries.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, accountresp)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
}

func (server *Server) GetAllAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       "Invalid input",
			"description": err.Error(),
		})
		return
	}
	args := db.ListAccountParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accountresp, err := server.store.ListAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, accountresp)

}

type getAccountID struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAccountByID(ctx *gin.Context) {
	var req getAccountID
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"description": "Invalid request",
			"error":       err.Error(),
		})
		return
	}

	respAccount, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
		return
	}
	ctx.JSON(http.StatusOK, respAccount)

}

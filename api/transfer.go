package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/shama3541/simplebank/db/database"
)

type TransferRequest struct {
	FromAccount int64  `json:"from_account" binding:"required,min=1"`
	ToAccount   int64  `json:"to_account" binding:"required,gt=0"`
	Amount      int64  `json:"amount" binding:"required,min=1"`
	Currency    string `json:"currency" binding:"required,oneof=USD EUR INR"`
}

func (server *Server) CreateTransfer(ctx *gin.Context) {
	var req TransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errmsg": "Invalid params or bad request",
			"error":  err.Error(),
		})
		return
	}
	if !server.ValidAccount(ctx, req.FromAccount, req.Currency) {
		return
	}

	if !server.ValidAccount(ctx, req.ToAccount, req.Currency) {
		return
	}

	args := db.TransferParams{
		FromAccountID: req.FromAccount,
		ToAccountID:   req.ToAccount,
		Amount:        req.Amount,
	}

	response, err := server.store.TranferTx(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) ValidAccount(ctx *gin.Context, accounID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accounID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "account not found",
			})
			return false
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return false
	}
	if currency != account.Currency {
		currencyerror := fmt.Errorf("Mismatch in currency type. Expected currency: %s Received Currency: %s", currency, account.Currency)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": currencyerror.Error(),
		})
		return false
	}
	return true
}

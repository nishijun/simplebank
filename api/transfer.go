package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nishijun/simplebank/db/sqlc"
)

type createTransferRequest struct {
	FromAccoundID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccoundID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if !server.validAccount(ctx, req.FromAccoundID, req.Currency) || !server.validAccount(ctx, req.ToAccoundID, req.Currency) {
		return
	}

	result, err := server.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccoundID,
		ToAccountID:   req.ToAccoundID,
		Amount:        req.Amount,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return false
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}

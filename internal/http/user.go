package http

import (
	"context"
	"net/http"

	"github.com/anti-duhring/go-grpc/internal/db"
	"github.com/anti-duhring/go-grpc/internal/errors"
	"github.com/anti-duhring/go-grpc/internal/grpc_client"
	"github.com/anti-duhring/go-grpc/internal/invoicer"
	"github.com/anti-duhring/go-grpc/internal/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (u *CreateUserRequest) Validate() error {
	if u.Name == "" {
		return errors.New(errors.ErrMissingRequiredField, "missing name field")
	}

	return nil
}

type TransferMoneyRequest struct {
	To     uuid.UUID `json:"to"`
	Amount float64   `json:"amount"`
}

func (t *TransferMoneyRequest) Validate() error {
	if t.To == uuid.Nil {
		return errors.New(errors.ErrMissingRequiredField, "missing to field")
	}

	if t.Amount <= 0 {
		return errors.New(errors.ErrMissingRequiredField, "missing amount field")
	}

	return nil
}

func GetUserBalance(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, "missing id parameter")
		return
	}

	user := schema.User{}

	if err := db.Client.Preload("Wallet").Where("id = ?", id).First(&user).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusOK, gin.H{"balance": user.Wallet.Balance})
}

func CreateUser(c *gin.Context) {
	request := CreateUserRequest{}

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	wallet := schema.Wallet{}

	if err := db.Client.Create(&wallet).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user := schema.User{
		Name:     request.Name,
		WalletID: wallet.ID,
	}

	if err := db.Client.Create(&user).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusCreated, user)
}

func TransferMoney(c *gin.Context) {
	from := c.Param("id")
	request := TransferMoneyRequest{}

	if from == "" {
		sendError(c, http.StatusBadRequest, "missing id parameter")
		return
	}

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	inv := invoicer.NewInvoicerClient(grpc_client.Conn)

	message := invoicer.CreateRequest{
		From: from,
		To:   request.To.String(),
		Amount: &invoicer.Amount{
			Amount:   int64(request.Amount),
			Currency: "USD",
		},
	}

	res, err := inv.Create(context.Background(), &message)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
	}

	sendSuccess(c, http.StatusOK, res)

}

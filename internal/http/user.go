package http

import (
	"context"
	"net/http"

	"github.com/anti-duhring/go-grpc/internal/db"
	"github.com/anti-duhring/go-grpc/internal/errors"
	"github.com/anti-duhring/go-grpc/internal/grpc_client"
	"github.com/anti-duhring/go-grpc/internal/schema"
	"github.com/anti-duhring/go-grpc/internal/wallet"
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

	userId := uuid.New()
	w := wallet.NewWalletServiceClient(grpc_client.Conn)

	message := wallet.CreateRequest{
		UserId: userId.String(),
	}

	res, err := w.Create(context.Background(), &message)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	walletId, err := uuid.Parse(res.Wallet.Id)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user := schema.User{
		Name:     request.Name,
		WalletID: walletId,
	}
	user.ID = userId

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

	w := wallet.NewWalletServiceClient(grpc_client.Conn)

	message := wallet.TransferRequest{
		From:   from,
		To:     request.To.String(),
		Amount: float32(request.Amount),
	}

	res, err := w.Transfer(context.Background(), &message)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusOK, res)

}

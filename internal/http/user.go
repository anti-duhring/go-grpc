package http

import (
	"net/http"

	"github.com/anti-duhring/go-grpc/internal/db"
	"github.com/anti-duhring/go-grpc/internal/errors"
	"github.com/anti-duhring/go-grpc/internal/schema"
	"github.com/gin-gonic/gin"
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

	sendSuccess(c, http.StatusOK, user)
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

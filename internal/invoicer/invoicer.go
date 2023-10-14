package invoicer

import (
	context "context"
	"log"

	"github.com/anti-duhring/go-grpc/internal/db"
	"github.com/anti-duhring/go-grpc/internal/errors"
	"github.com/anti-duhring/go-grpc/internal/schema"
)

type Server struct {
	UnimplementedInvoicerServer
}

func (s *Server) Create(crx context.Context, req *CreateRequest) (*CreateResponse, error) {
	log.Printf("Received message from client")

	err := db.Initialize()

	if err != nil {
		return nil, err
	}

	from := schema.Wallet{}
	to := schema.Wallet{}

	if err := db.Client.Where("id = ?", req.From).First(&from).Error; err != nil {
		return nil, err
	}

	if from.Balance < float64(req.Amount.Amount) {
		return nil, errors.New(errors.ErrWalletDontHaveEnoughMoney, "Wallet dont have enough money")
	}

	if err := db.Client.Where("id = ?", req.To).First(&to).Error; err != nil {
		return nil, err
	}

	from.Balance -= float64(req.Amount.Amount)
	to.Balance += float64(req.Amount.Amount)

	if err := db.Client.Save(&from).Error; err != nil {
		return nil, err
	}

	if err := db.Client.Save(&to).Error; err != nil {
		return nil, err
	}

	return &CreateResponse{
		FromAmount: int64(from.Balance),
		ToAmount:   int64(to.Balance),
	}, nil
}

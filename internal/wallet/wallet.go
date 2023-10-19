package wallet

import (
	"context"
	"os"

	"github.com/anti-duhring/go-grpc/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	UnimplementedWalletServiceServer
}

func (s *Server) Create(c context.Context, request *CreateRequest) (*CreateResponse, error) {

	db, err := dbConn()

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	defer sqlDB.Close()

	wallet := schema.Wallet{}

	if err := db.Table(wallet.TableName()).Create(&wallet).Error; err != nil {
		return nil, err
	}

	return &CreateResponse{
		UserId: request.UserId,
		Wallet: &Wallet{
			Id:       wallet.ID.String(),
			Balance:  float32(wallet.Balance),
			Currency: wallet.Currency,
		},
	}, nil
}

func (s *Server) Transfer(c context.Context, request *TransferRequest) (*TransferResponse, error) {

	db, err := dbConn()

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	defer sqlDB.Close()

	from := schema.Wallet{}
	to := schema.Wallet{}

	if err := db.Table(from.TableName()).Where("id = ?", request.From).First(&from).Error; err != nil {
		return nil, err
	}

	if err := db.Table(to.TableName()).Where("id = ?", request.To).First(&to).Error; err != nil {
		return nil, err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		var err error

		if err = transferMoney(&from, &to, float64(request.Amount)); err != nil {
			return err
		}

		if err = tx.Table(from.TableName()).Save(&from).Error; err != nil {
			return err
		}

		if err = tx.Table(to.TableName()).Save(&to).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &TransferResponse{
		Status: "success",
	}, nil
}

func dbConn() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	return client, err
}

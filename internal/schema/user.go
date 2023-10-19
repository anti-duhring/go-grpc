package schema

import (
	"github.com/google/uuid"
)

type User struct {
	Base
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	WalletID uuid.UUID `gorm:"type:uuid;not null" json:"wallet_id"`
	Wallet   Wallet    `gorm:"foreignKey:WalletID" json:"wallet"`
}

func (u *User) TableName() string {
	return "users"
}

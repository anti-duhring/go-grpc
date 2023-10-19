package schema

type Wallet struct {
	Base
	Balance  float64 `gorm:"not null;default:0" json:"balance"`
	Currency string  `gorm:"not null;default:USD" json:"currency"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}

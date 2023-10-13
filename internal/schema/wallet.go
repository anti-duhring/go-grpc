package schema

type Wallet struct {
	Base
	Balance float64 `gorm:"not null;default:0" json:"balance"`
}

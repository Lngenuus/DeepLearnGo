package orm

import (
	"gorm.io/gorm"
)

type Coin struct {
	gorm.Model
	ID     int64 `gorm:"primaryKey" json:"id"`
	Amount int64 `json:"amount"`
}

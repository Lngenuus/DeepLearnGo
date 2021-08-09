package data

import (
	"coin/internal/conf"
	"coin/internal/data/orm"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCoinRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	localDB := "../../internal/data/test.db"
	log := log.NewHelper(logger)
	db, err := gorm.Open(sqlite.Open(localDB), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	db.AutoMigrate(&orm.Coin{})
	db.Create(&orm.Coin{ID: 1, Amount: 0})

	d := &Data{db: db}

	return d, func() {
		log.Info("message", "closing the data resources")
		os.Remove(localDB)
	}, nil
}

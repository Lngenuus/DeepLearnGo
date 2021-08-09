package data

import (
	"coin/internal/biz"
	"coin/internal/data/orm"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type coinRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewCoinRepo(data *Data, logger log.Logger) biz.CoinRepo {
	return &coinRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *coinRepo) SetCoin(ctx context.Context, g *biz.Coin) error {
	tx := r.data.db.Begin()
	tx.Model(&orm.Coin{}).Where("id", 1).Update("amount", g.Num)
	tx.Commit()
	return nil
}

func (r *coinRepo) GetCoin(ctx context.Context) (*biz.Coin, error) {
	var c orm.Coin
	r.data.db.Find(&c, 1)
	return &biz.Coin{Num: c.Amount}, nil
}

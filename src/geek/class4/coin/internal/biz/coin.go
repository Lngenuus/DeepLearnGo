package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Coin struct {
	Num int64
}

type CoinRepo interface {
	SetCoin(context.Context, *Coin) error
	GetCoin(context.Context) (*Coin, error)
}

type CoinOptcase struct {
	repo CoinRepo
	log  *log.Helper
}

func NewCoinOptcase(repo CoinRepo, logger log.Logger) *CoinOptcase {
	return &CoinOptcase{repo: repo, log: log.NewHelper(logger)}
}

func (op *CoinOptcase) AddCoin(ctx context.Context, n int64) error {
	oldCoin, err := op.repo.GetCoin(ctx)
	if err != nil {
		return err
	}
	return op.repo.SetCoin(ctx, &Coin{Num: oldCoin.Num + n})
}

func (op *CoinOptcase) ReduceCoin(ctx context.Context, n int64) error {
	oldCoin, err := op.repo.GetCoin(ctx)
	if err != nil {
		return err
	}
	return op.repo.SetCoin(ctx, &Coin{Num: oldCoin.Num - n})
}

func (op *CoinOptcase) ShowCoin(ctx context.Context) (*Coin, error) {
	oldCoin, err := op.repo.GetCoin(ctx)
	if err != nil {
		return nil, err
	}
	return oldCoin, nil
}

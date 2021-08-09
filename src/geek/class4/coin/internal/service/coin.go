package service

import (
	"coin/internal/biz"
	"context"

	pb "coin/api/coin/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type CoinService struct {
	pb.UnimplementedCoinServer

	op  *biz.CoinOptcase
	log *log.Helper
}

func NewCoinService(op *biz.CoinOptcase, logger log.Logger) *CoinService {
	return &CoinService{op: op, log: log.NewHelper(logger)}
}

func (s *CoinService) AddCoin(ctx context.Context, req *pb.AddCoinRequest) (*pb.AddCoinReply, error) {
	if err := s.op.AddCoin(ctx, req.GetNum()); err != nil {
		return nil, err
	}
	return &pb.AddCoinReply{Message: "success"}, nil
}
func (s *CoinService) ReduceCoin(ctx context.Context, req *pb.ReduceCoinRequest) (*pb.ReduceCoinReply, error) {
	if err := s.op.ReduceCoin(ctx, req.GetNum()); err != nil {
		return nil, err
	}
	return &pb.ReduceCoinReply{Message: "success"}, nil
}
func (s *CoinService) ShowCoin(ctx context.Context, req *pb.ShowCoinRequest) (*pb.ShowCoinReply, error) {
	tCoin, err := s.op.ShowCoin(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ShowCoinReply{Amount: tCoin.Num, Message: "success"}, nil
}

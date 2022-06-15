package service

import (
	"context"
	"github.com/prometheus/common/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"load_balance/server/proto"
)

type Service struct{}

func (s Service) Test(ctx context.Context, empty *emptypb.Empty) (*proto.TestResponse, error) {
	log.Info("收到一个请求\n")
	return &proto.TestResponse{Msg: "test"}, nil
}

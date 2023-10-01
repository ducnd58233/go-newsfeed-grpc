package service

import (
	"github.com/ducnd58233/go-newsfeed-grpc/app/configs"
	"github.com/ducnd58233/go-newsfeed-grpc/app/internal/logger"
	auth_client "github.com/ducnd58233/go-newsfeed-grpc/app/pkg/client/auth"
	"github.com/ducnd58233/go-newsfeed-grpc/app/pkg/types/proto/pb/auth_pb"
	"go.uber.org/zap"
)

type Service struct {
	authClient auth_pb.AuthClient
	log *zap.Logger
}

func NewService(conf *configs.AppConfig) (*Service, error) {
	log, err := logger.NewLogger()
	if err != nil {
		log.Error("failed to init logger", zap.Error(err))
		return nil, err
	}

	authClient, err := auth_client.NewClient(conf.Authentication.Hosts)
	if err != nil {
		log.Error("failed to init auth client", zap.Error(err))
		return nil, err
	}

	return &Service{
		authClient: authClient,
		log: log,
	}, nil
}
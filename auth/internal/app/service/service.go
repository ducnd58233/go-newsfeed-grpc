package service

import (
	"context"
	"time"

	"github.com/ducnd58233/go-newsfeed-grpc/auth/configs"
	"github.com/ducnd58233/go-newsfeed-grpc/auth/internal/logger"
	"github.com/ducnd58233/go-newsfeed-grpc/auth/pkg/types/proto/pb/auth_pb"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var log *zap.Logger

type AuthService struct {
	auth_pb.UnimplementedAuthServer
	db    *gorm.DB
	redis *redis.Client

	log *zap.Logger
}

func (a *AuthService) CheckUserAuthentication(ctx context.Context, info *auth_pb.CheckUserAuthenticationRequest) (*auth_pb.CheckUserAuthenticationResponse, error) {
	panic("implement me")
}

func (a *AuthService) CreateUser(ctx context.Context, info *auth_pb.UserDetailInfo) (*auth_pb.UserResult, error) {
	panic("implement me")
}

func (a *AuthService) EditUser(ctx context.Context, info *auth_pb.EditUserRequest) (*auth_pb.EditUserResponse, error) {
	panic("implement me")
}

func (a *AuthService) GetUserFollower(ctx context.Context, info *auth_pb.UserInfo) (*auth_pb.UserFollower, error) {
	panic("implement me")
}

func (a *AuthService) FollowUser(ctx context.Context, info *auth_pb.FollowUserRequest) (*auth_pb.FollowUserResponse, error) {
	panic("implement me")
}

func (a *AuthService) UnfollowUser(ctx context.Context, info *auth_pb.UnfollowUserRequest) (*auth_pb.UnfollowUserResponse, error) {
	panic("implement me")
}

func (a *AuthService) GetFollowerList(ctx context.Context, info *auth_pb.GetFollowerListRequest) (*auth_pb.GetFollowerListResponse, error) {
	panic("implement me")
}

func NewAuthService(conf *configs.AuthConfig) (*AuthService, error) {
	var err error
	log, err = logger.NewLogger()
	if err != nil {
		log.Error("failed to init logger", zap.Error(err))
		return nil, err
	}

	db := connectToDB(conf.MySQL)

	rd := connectToRedis(&conf.Redis)

	return &AuthService{db: db, redis: rd, log: log}, nil
}

func connectToDB(db mysql.Config) *gorm.DB {
	var counts int

	for {
		conn, err := gorm.Open(mysql.New(db), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			log.Error("MySQL not yet ready ...")
			counts++
		} else {
			log.Info("Connected to MySQL")
			return conn
		}

		if counts > 10 {
			log.Error("Error while connecting MySQL", zap.Error(err))
			return nil
		}

		log.Info("Backing off for 2 seconds ...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func connectToRedis(db *redis.Options) *redis.Client {
	var counts int64

	for {
		rd := redis.NewClient(db)
		if rd == nil {
			log.Error("Redis not yet ready ...")
			counts++
		} else {
			log.Info("Connected to MySQL")
			return rd
		}

		if counts > 10 {
			log.Error("Error while connecting Redis")
			return nil
		}

		log.Info("Backing off for 2 seconds ...")
		time.Sleep(2 * time.Second)
		continue
	}
}

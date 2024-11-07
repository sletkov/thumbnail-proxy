package main

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/redis/go-redis/v9"
	"github.com/sletkov/thumbnail-proxy/config"
	"github.com/sletkov/thumbnail-proxy/internal/cache"
	"github.com/sletkov/thumbnail-proxy/internal/service"
	"github.com/sletkov/thumbnail-proxy/internal/transport/grpcserver"
	"github.com/sletkov/thumbnail-proxy/internal/transport/youtubeclient"
	proto "github.com/sletkov/thumbnail-proxy/pkg/sdk/go/thumbnailproxy_grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

const serviceID = "thumbnail-proxy"

var envPath string

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()).With(zap.String("service_id", serviceID)))
	flag.StringVar(&envPath, "env-path", ".env", "path to env file")
}

func main() {
	var cfg config.Config
	flag.Parse()
	cleanenv.ReadConfig(envPath, &cfg)
	zap.L().Info(fmt.Sprintf("service started"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})

	cache := cache.New(rdb, cfg.RedisTTL)
	transport := youtubeclient.New(cfg.YoutubeAPIKey)
	thumbnailService := service.New(cache, transport)
	grpcServer := grpcserver.New(thumbnailService)
	srv := grpc.NewServer()
	proto.RegisterThumbnailProxyServer(srv, grpcServer)

	lis, err := net.Listen("tcp", net.JoinHostPort(cfg.GRPCServerHost, cfg.GRPCServerPort))
	if err != nil {
		zap.L().Fatal("failed to listen tcp", zap.Error(err))
	}

	zap.L().Info("grpc server started", zap.String("address:", net.JoinHostPort(cfg.GRPCServerHost, cfg.GRPCServerPort)))
	if err := srv.Serve(lis); err != nil {
		zap.L().Fatal("service stopped", zap.Error(err))
	}
}

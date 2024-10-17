package config

import "time"

type Config struct {
	GRPCServerHost string        `env:"GRPC_SERVER_HOST" env-default:"localhost"`
	GRPCServerPort string        `env:"GRPC_SERVER_PORT" env-default:"8083"`
	RedisHost      string        `env:"REDIS_HOST" env-default:"localhost"`
	RedisPort      string        `env:"REDIS_PORT" env-default:"6379"`
	RedisTTL       time.Duration `env:"REDIS_TTL" env-default:"24h"`
	RedisPassword  string        `env:"REDIS_PASSWORD" env-default:""`
	YoutubeAPIKey  string        `env:"YOUTUBE_API_KEY" env-required:"true"`
}

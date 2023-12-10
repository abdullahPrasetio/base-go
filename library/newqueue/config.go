package newqueue

import (
	"strconv"

	"github.com/abdullahPrasetio/base-go/configs"
	"github.com/hibiken/asynq"
)

func NewRedisConfig() asynq.RedisClientOpt {
	config := configs.Configs
	db, _ := strconv.Atoi(config.REDIS_DB_QUEUE)
	return asynq.RedisClientOpt{
		Addr: config.REDIS_HOST,
		DB:   db, Password: config.REDIS_PASSWORD,
	}
}

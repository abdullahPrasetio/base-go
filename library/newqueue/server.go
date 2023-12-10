package newqueue

import (
	"github.com/hibiken/asynq"
)

func NewServer() *asynq.Server {
	srv := asynq.NewServer(NewRedisConfig(),
		asynq.Config{Concurrency: 10, Queues: map[string]int{
			"critical": 6,
			"default":  3}},
	)
	return srv
}

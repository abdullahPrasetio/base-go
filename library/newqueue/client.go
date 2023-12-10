package newqueue

import (
	"github.com/hibiken/asynq"
)

func NewClient() *asynq.Client {

	client := asynq.NewClient(NewRedisConfig())

	return client
}

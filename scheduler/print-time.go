package scheduler

import (
	"encoding/json"
	"fmt"

	"github.com/abdullahPrasetio/base-go/library/newqueue"
	"github.com/abdullahPrasetio/base-go/utils/times"
	"github.com/hibiken/asynq"
)

func NewPrintTimeScheduller() {

	// Buat scheduler dengan store yang sama dengan server
	scheduler := asynq.NewScheduler(newqueue.NewRedisConfig(), nil)
	payload, err := json.Marshal(times.TimeNow{ValueTime: times.GetTimeNow("").ValueTime, ValueString: times.GetTimeNow("").ValueString})
	if err != nil {
		fmt.Println("err", err)
	}
	// crontab by unix time
	if _, err := scheduler.Register("@every 30s", asynq.NewTask("printTime", payload)); err != nil {
		fmt.Println("err", err)
	}
	if err := scheduler.Start(); err != nil {
		fmt.Println("err", err)
	}

}

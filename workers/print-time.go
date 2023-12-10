package workers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/abdullahPrasetio/base-go/library/newqueue"
	"github.com/abdullahPrasetio/base-go/utils/times"
	"github.com/hibiken/asynq"
)

func ProcessTimeScheduller() {
	server := newqueue.NewServer()

	mux := asynq.NewServeMux()
	mux.HandleFunc("printTime", PrintTimeProcess)

	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}

}

func PrintTimeProcess(ctx context.Context, t *asynq.Task) error {

	var p times.TimeNow
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Successfully proccess Task task: %+v", p.ValueString)
	return nil
}

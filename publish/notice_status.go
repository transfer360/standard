package publish

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

func PublishNoticeStatusChange(ctx context.Context, sref string, statusCode int, source string) {

	// Creates a client.
	client, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Errorln("Failed to create client: %v", err)
		return
	}

	msg, _ := json.Marshal(struct {
		Sref       string `json:"sref"`
		StatusCode int    `json:"status_code"`
	}{
		Sref:       sref,
		StatusCode: statusCode,
	})

	t := client.Topic("notice_status_change")
	_ = t.Publish(ctx, &pubsub.Message{Data: msg, Attributes: map[string]string{"source": source}, OrderingKey: "status_update"})

}

package publish

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func StatusUpdate(ctx context.Context, sref string, statusCode int, statusDescription string, source string) {

	// Creates a client.
	client, err := pubsub.NewClient(ctx, "transfer-360")
	if err != nil {
		log.Errorln("Failed to create client:", err)
		return
	}

	// Sets the id of the topic.
	topic := client.Topic("notice_status_update_webhook")

	jstr, err := json.Marshal(struct {
		Sref              string `json:"sref"`
		StatusCode        int    `json:"status_code"`
		StatusDescription string `json:"status_description"`
	}{
		Sref:              sref,
		StatusCode:        statusCode,
		StatusDescription: statusDescription,
	})

	if err != nil {
		log.Errorln(err)
		return
	}

	// Sends a message.
	result := topic.Publish(ctx, &pubsub.Message{
		Data: jstr,
		Attributes: map[string]string{
			"source": source,
		},
	})

	_, err = result.Get(ctx)
	if err != nil {
		log.Errorln(err)
	}

}

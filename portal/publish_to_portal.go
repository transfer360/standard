package portal

import (
	"cloud.google.com/go/pubsub"
	"context"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// SendEvent ------------------------------------------------------------------------------------------
func SendEvent(ctx context.Context, data []byte, event string, attr map[string]string, ClientID string, ProjectID string) {

	if attr == nil {
		attr = make(map[string]string)
	}

	client, err := pubsub.NewClient(ctx, ProjectID)
	if err != nil {
		log.Errorf("Unable to publish results: %v", err)
		return
	}

	topicName := "save_to_portal"
	topic := client.Topic(topicName)

	attr["event"] = event
	attr["client"] = ClientID
	attr["date_published"] = time.Now().Format(time.RFC3339)

	msg := &pubsub.Message{
		Data:       data,
		Attributes: attr,
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Errorf("Unable to publish results: %v", err)
		return
	}

}

// SendEventSandBox ------------------------------------------------------------------------------------------------------------------------------------
func SendEventSandBox(ctx context.Context, data []byte, event string, attr map[string]string, ClientID string, ProjectID string) {

	if attr == nil {
		attr = make(map[string]string)
	}

	attr["sandbox"] = strconv.FormatBool(true)

	SendEvent(ctx, data, event, attr, ClientID, ProjectID)
}

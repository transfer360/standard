package status

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/transfer360/standard/database"
	"github.com/transfer360/standard/publish"
	"google.golang.org/api/iterator"
	"time"
)

var ErrSREFNotFound = errors.New("sref not found")

func SetStatus(ctx context.Context, SRef string, statusCode int, statusDescription string, source string) error {

	fsclient, err := database.GetFirebaseClient(ctx, "transfer-360")
	if err != nil {
		log.Errorln("WillPay:", err)
		return err

	}

	srefFound := false

	itr := fsclient.Collection("searches").Where("sref", "==", SRef).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				break
			} else {
				return err
			}
		} else {
			if doc.Exists() {
				srefFound = true
				_, err = doc.Ref.Update(ctx, []firestore.Update{
					{
						Path:  "status",
						Value: statusCode,
					},
					{
						Path:  "status_description",
						Value: statusDescription,
					},
					{
						Path:  "status_changed",
						Value: time.Now(),
					},
				})

				if err != nil {
					return err
				}
			}
		}
	}

	if !srefFound {
		return ErrSREFNotFound
	} else {
		publish.StatusUpdate(ctx, SRef, statusCode, statusDescription, source)
		publish.PublishNoticeStatusChange(ctx, SRef, statusCode, statusDescription)
	}

	return nil

}

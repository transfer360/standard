package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// GetFirebaseClient ---------------------------------------------------------------------------------------------
func GetFirebaseClient(ctx context.Context, projectID string) (*firestore.Client, error) {

	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
		return nil, fmt.Errorf("unexpected error creating firebase client")
	}

	return app.Firestore(ctx)

}

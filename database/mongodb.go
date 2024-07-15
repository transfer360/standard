package database

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type creds struct {
	Private  string `json:"private"`
	Public   string `json:"public"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func MongoDB(ctx context.Context) (*mongo.Client, error) {

	dbc, err := getMongoDBCredentials(ctx)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	connectionURL := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=t360", dbc.Username, dbc.Password, dbc.Private)
	if len(os.Getenv("DEVELOPMENT")) > 0 {
		connectionURL = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=t360", dbc.Username, dbc.Password, dbc.Public)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionURL).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Errorln(err)
	}

	return client, nil

}

// getCredentials ------------------------------------------------------------------------------------------
func getMongoDBCredentials(ctx context.Context) (creds, error) {

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Println(fmt.Errorf("failed to create secretmanager client: %v", err))
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/569859308728/secrets/mongodb/versions/latest",
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Println(fmt.Errorf("failed to access secret version: %v", err))
	}

	dbc := creds{}

	err = json.Unmarshal(result.Payload.Data, &dbc)

	return dbc, err

}

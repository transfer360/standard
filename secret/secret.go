package secret

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func Get(ctx context.Context, path string) ([]byte, error) {

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Errorf("CredentialsFromSecretManagerPath:%v", err)
		return []byte{}, fmt.Errorf("failed to create secretmanager client: %w", err)
	}

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: path,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Errorf("CredentialsFromSecretManagerPath:%v", err)
		return []byte{}, fmt.Errorf("failed to get secret version: %w", err)
	}

	return result.Payload.Data, nil

}

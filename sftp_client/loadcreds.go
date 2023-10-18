package sftp_client

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type ConnectionDetails struct {
	Host     string `json:"host"`
	Username string `json:"user"`
	Password string `json:"password"`
}

func LoadSFTPSecrets(ctx context.Context, secretPath string) (sfd ConnectionDetails, err error) {

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Errorln(fmt.Errorf("failed to create secretmanager client: %v", err))
		return sfd, err
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretPath,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Println(fmt.Errorf("failed to access secret version: %v", err))
		return sfd, err
	}

	err = json.Unmarshal(result.Payload.Data, &sfd)

	return sfd, err

}

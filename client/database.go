package client

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/transfer360/standard/database"
)

var ErrClientSecretCredentialsNotFound = errors.New("secret credentials not found")

func GetDatabaseCredentials(ctx context.Context, clientID string, db *sql.DB) (dbc database.DatabaseConfiguration, err error) {

	var clientSecretPath string

	err = db.QueryRow("SELECT secret_path FROM portal_access WHERE t360_id=?", clientID).Scan(&clientSecretPath)
	if errors.Is(err, sql.ErrNoRows) {
		return dbc, fmt.Errorf("%w for %s", ErrClientSecretCredentialsNotFound, clientID)
	}

	return database.CredentialsFromSecretManagerPath(ctx, clientSecretPath)
}

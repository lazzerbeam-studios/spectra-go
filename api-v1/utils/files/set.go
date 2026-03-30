package files

import (
	"context"
	"encoding/base64"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const UrlGCP = "https://storage.googleapis.com/"

var StorageClientGCP *StorageClientServiceGCP

type StorageClientServiceGCP struct {
	bucket        string
	project       string
	storageClient *storage.Client
}

func SetClientGCP(credentials string, project string, bucket string) {
	credentialsByte, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		panic("failed to decode GCP credentials")
	}

	storageClient, err := storage.NewClient(
		context.Background(),
		option.WithCredentialsJSON(credentialsByte),
	)
	if err != nil {
		panic("failed to create GCP storage client")
	}

	StorageClientGCP = &StorageClientServiceGCP{
		bucket:        bucket,
		project:       project,
		storageClient: storageClient,
	}
}

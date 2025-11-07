package files

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const UrlGCP = "https://storage.googleapis.com/"

var ClientGCP *StorageClientGCP

type StorageClientGCP struct {
	bucket        string
	project       string
	storageClient *storage.Client
}

func (client *StorageClientGCP) UploadFile(path string, content []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	writer := client.storageClient.Bucket(client.bucket).Object(path).NewWriter(ctx)
	reader := bytes.NewBuffer(content)

	if _, err := io.Copy(writer, reader); err != nil {
		return "", err
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	url := UrlGCP + client.bucket + "/" + path
	return url, nil
}

func SetStorageGCP(credentials string, project string, bucket string) {
	credentialsByte, err := base64.StdEncoding.DecodeString(credentials)
	if err != nil {
		panic("Failed to decode GCP credentials")
	}

	storageClient, err := storage.NewClient(
		context.Background(),
		option.WithCredentialsJSON(credentialsByte),
	)
	if err != nil {
		panic("Failed to create GCP storage client")
	}

	ClientGCP = &StorageClientGCP{
		bucket:        bucket,
		project:       project,
		storageClient: storageClient,
	}
}

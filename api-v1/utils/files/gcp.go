package files

import (
	"bytes"
	"context"
	"io"
	"time"
)

func (client *StorageClientServiceGCP) UploadFile(path string, content []byte) (string, error) {
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

func (client *StorageClientServiceGCP) DownloadFile(path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	reader, err := client.storageClient.Bucket(client.bucket).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)

	return data, err
}

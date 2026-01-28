package files

import (
	"bytes"
	"context"
	"io"
	"time"
)

func (client *StorageClientServiceGCP) UploadFileSVG(path string, content []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	writer := client.storageClient.Bucket(client.bucket).Object(path).NewWriter(ctx)
	writer.ContentType = "image/svg+xml"
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

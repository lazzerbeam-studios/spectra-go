package bunny

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

const defaultContentType = "application/octet-stream"

func UploadMedia(objectPath string, fileData []byte, contentType string) (string, error) {
	if Storage == nil {
		return "", errors.New("bunny storage client not initialized")
	}
	if Storage.storageZone == "" {
		return "", errors.New("bunny storage zone not configured")
	}
	if objectPath == "" {
		return "", errors.New("object path is required")
	}

	normalizedPath := strings.TrimPrefix(objectPath, "/")
	uploadURL := buildUploadURL(normalizedPath)
	request, err := http.NewRequest("PUT", uploadURL, bytes.NewReader(fileData))
	if err != nil {
		return "", errors.New("error creating upload request")
	}

	request.Header.Set("AccessKey", Storage.apiKey)
	if contentType == "" {
		request.Header.Set("Content-Type", defaultContentType)
	} else {
		request.Header.Set("Content-Type", contentType)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", errors.New("error making upload request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(response.Body)
		return "", errors.New("failed to upload file. Status: " + response.Status + ", Body: " + string(bodyBytes))
	}

	return buildPublicURL(normalizedPath), nil
}

func buildUploadURL(normalizedPath string) string {
	storageHost := "storage.bunnycdn.com"
	if Storage.region != "" {
		storageHost = Storage.region + ".storage.bunnycdn.com"
	}
	return "https://" + storageHost + "/" + Storage.storageZone + "/" + normalizedPath
}

func buildPublicURL(normalizedPath string) string {
	if Storage.cdnHost == "" {
		return normalizedPath
	}

	cdnBaseURL := strings.TrimSuffix(strings.TrimSpace(Storage.cdnHost), "/")
	if !strings.HasPrefix(cdnBaseURL, "http://") && !strings.HasPrefix(cdnBaseURL, "https://") {
		cdnBaseURL = "https://" + cdnBaseURL
	}

	return cdnBaseURL + "/" + normalizedPath
}

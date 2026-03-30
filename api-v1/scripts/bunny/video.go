package bunny

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const bunnyURL = "https://video.bunnycdn.com"

type VideoResponse struct {
	VideoID string `json:"guid"`
	Title   string `json:"title"`
}

func CreateVideoObject(title string) (string, error) {
	videoPayload := map[string]string{"title": title}
	videoJsonPayload, err := json.Marshal(videoPayload)
	if err != nil {
		return "", errors.New("error marshalling json")
	}

	url := bunnyURL + "/library/" + Video.libraryID + "/videos"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(videoJsonPayload))
	if err != nil {
		return "", errors.New("error creating request")
	}

	request.Header.Set("AccessKey", Video.apiKey)
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", errors.New("error making request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return "", errors.New("failed to create video object. Status: " + response.Status + ", Body: " + string(bodyBytes))
	}

	var videoResponse VideoResponse
	if err := json.NewDecoder(response.Body).Decode(&videoResponse); err != nil {
		return "", errors.New("error decoding response")
	}

	return videoResponse.VideoID, nil
}

func UploadVideoFile(videoID string, videoData []byte) error {
	url := bunnyURL + "/library/" + Video.libraryID + "/videos/" + videoID
	request, err := http.NewRequest("PUT", url, bytes.NewReader(videoData))
	if err != nil {
		return errors.New("error creating upload request")
	}

	request.Header.Set("AccessKey", Video.apiKey)
	request.Header.Set("Content-Type", "application/octet-stream")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.New("error making upload request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return errors.New("failed to upload video file. Status: " + response.Status + ", Body: " + string(bodyBytes))
	}

	return nil
}

func UploadVideo(videoData []byte, title string) (string, error) {
	if Video == nil {
		return "", errors.New("bunny client not initialized")
	}

	videoID, err := CreateVideoObject(title)
	if err != nil {
		return "", err
	}

	if err := UploadVideoFile(videoID, videoData); err != nil {
		return "", err
	}

	return videoID, nil
}

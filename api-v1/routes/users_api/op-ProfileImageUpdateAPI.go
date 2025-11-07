package users_api

import (
	"context"
	"io"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/models"
	"api-go/utils/auth"
	"api-go/utils/files"
)

type ProfileImageUpdateInput struct {
	auth.AuthParam
	RawBody huma.MultipartFormFiles[struct {
		Image huma.FormFile `form:"file" contentType:"text/plain" required:"true"`
	}]
}

func (input *ProfileImageUpdateInput) Resolve(ctx huma.Context) []error {
	var err error
	input.User, err = auth.AuthUser(input.Auth)
	if err != nil {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	return nil
}

type ProfileImageUpdateOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileImageUpdateAPI(ctx context.Context, input *ProfileImageUpdateInput) (*ProfileImageUpdateOutput, error) {
	imageData := input.RawBody.Data()
	image, err := io.ReadAll(imageData.Image)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to read image.")
	}

	fileDir := "users/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-avatar.png"
	fileUrl, err := files.ClientGCP.UploadFile(fileDir, image)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to upload image.")
	}

	profileObj, err := input.User.Update().
		SetImage(fileUrl).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to update user.")
	}

	response := &ProfileImageUpdateOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}

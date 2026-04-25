package users_api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/models"
	"api-go/utils/auth"
)

type ProfileDeleteInput struct {
	auth.AuthParam
}

func (input *ProfileDeleteInput) Resolve(ctx huma.Context) []error {
	var err error
	input.User, err = auth.AuthUser(input.Auth)
	if err != nil {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	return nil
}

type ProfileDeleteOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileDeleteAPI(ctx context.Context, input *ProfileDeleteInput) (*ProfileDeleteOutput, error) {
	profileObj, err := input.User.Update().
		SetDeleted(true).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to delete user.")
	}

	response := &ProfileDeleteOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}

package users_api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/models"
	"api-go/utils/auth"
)

type ProfileUpdateInput struct {
	auth.AuthParam
	Body struct {
		models.ProfileUpdate
	}
}

func (input *ProfileUpdateInput) Resolve(ctx huma.Context) []error {
	var err error
	input.User, err = auth.AuthUser(input.Auth)
	if err != nil {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	return nil
}

type ProfileUpdateOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileUpdateAPI(ctx context.Context, input *ProfileUpdateInput) (*ProfileUpdateOutput, error) {
	profileObj, err := input.User.Update().
		SetName(input.Body.Name).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to update user.")
	}

	response := &ProfileUpdateOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}

package users_api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/models"
	"api-go/utils/auth"
)

type ProfileGetInput struct {
	auth.AuthParam
}

func (input *ProfileGetInput) Resolve(ctx huma.Context) []error {
	var err error
	input.User, err = auth.AuthUser(input.Auth)
	if err != nil {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	return nil
}

type ProfileGetOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileGetAPI(ctx context.Context, input *ProfileGetInput) (*ProfileGetOutput, error) {
	response := &ProfileGetOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &input.User)
	response.Body.Object = *object
	return response, nil
}

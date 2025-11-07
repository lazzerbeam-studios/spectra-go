package users_api

import (
	"context"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/ent/user"
	"api-go/models"
	"api-go/utils/auth"
	"api-go/utils/db"
)

type ProfileUpdateEmailInput struct {
	auth.AuthParam
	Body struct {
		models.ProfileUpdateEmail
	}
}

func (input *ProfileUpdateEmailInput) Resolve(ctx huma.Context) []error {
	var err error
	input.User, err = auth.AuthUser(input.Auth)
	if err != nil {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}

	input.Body.Email = strings.ToLower(input.Body.Email)
	usersList, err := db.EntDB.User.Query().Where(user.EmailEqualFold(input.Body.Email)).All(context.Background())
	if len(usersList) > 0 || err != nil {
		return []error{huma.Error400BadRequest("This email is not available.")}
	}

	return nil
}

type ProfileUpdateEmailOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileUpdateEmailAPI(ctx context.Context, input *ProfileUpdateEmailInput) (*ProfileUpdateEmailOutput, error) {
	profileObj, err := input.User.Update().
		SetEmail(input.Body.Email).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to update user.")
	}

	response := &ProfileUpdateEmailOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}

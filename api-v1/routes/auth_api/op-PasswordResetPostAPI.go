package auth_api

import (
	"context"
	"strconv"

	"github.com/danielgtaylor/huma/v2"

	"api-go/ent/user"
	"api-go/utils/cache"
	"api-go/utils/db"
	"api-go/utils/password"
)

type PasswordResetPostInput struct {
	Body struct {
		Code     string `json:"code"`
		Password string `json:"password"`
	}
}

func (input *PasswordResetPostInput) Resolve(ctx huma.Context) []error {
	var err error
	input.Body.Password, err = password.HashPassword(input.Body.Password)
	if err != nil {
		return []error{huma.Error400BadRequest("Unable to update password.")}
	}
	return nil
}

type PasswordResetPostOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

func PasswordResetPostAPI(ctx context.Context, input *PasswordResetPostInput) (*PasswordResetPostOutput, error) {
	userIDCache, err := cache.GetKey(input.Body.Code)
	if err != nil {
		return nil, huma.Error404NotFound("Code not found.")
	}

	userID, err := strconv.Atoi(userIDCache)
	if err != nil {
		return nil, huma.Error404NotFound("Code not found.")
	}

	userObj, err := db.EntDB.User.Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("User not found.")
	}

	_, err = userObj.Update().
		SetPassword(input.Body.Password).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to update password.")
	}

	response := &PasswordResetPostOutput{}
	response.Body.Message = "Your password has been changed successfully."
	return response, nil
}

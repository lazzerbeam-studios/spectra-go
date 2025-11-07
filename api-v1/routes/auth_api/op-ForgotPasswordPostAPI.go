package auth_api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/danielgtaylor/huma/v2"

	"api-go/ent/user"
	"api-go/scripts/generator"
	"api-go/utils/cache"
	"api-go/utils/db"
)

type ForgotPasswordPostInput struct {
	Body struct {
		Email string `json:"email"`
	}
}

type ForgotPasswordPostOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

func ForgotPasswordPostAPI(ctx context.Context, input *ForgotPasswordPostInput) (*ForgotPasswordPostOutput, error) {
	userObj, err := db.EntDB.User.Query().
		Where(user.EmailEqualFold(input.Body.Email)).
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("User not found.")
	}

	code := generator.GenerateRandomLetters() + strconv.Itoa(userObj.ID)
	cache.SetKey(code, strconv.Itoa(userObj.ID), 21600)

	fmt.Println(code)

	response := &ForgotPasswordPostOutput{}
	response.Body.Message = "Code sent to email."
	return response, nil
}

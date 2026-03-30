package auth_api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api-go/ent"
	"api-go/ent/user"
	"api-go/utils/auth"
	"api-go/utils/db"
	"api-go/utils/password"
)

type SignInSupabaseInput struct {
	Body struct {
		Token string `json:"token"`
	}
}

type SignInSupabaseOutput struct {
	Body struct {
		Token string `json:"token"`
	}
}

func SignInSupabaseAPI(ctx context.Context, input *SignInSupabaseInput) (*SignInSupabaseOutput, error) {
	tokenIn := strings.TrimSpace(input.Body.Token)
	if tokenIn == "" {
		return nil, huma.Error400BadRequest("token is required.")
	}

	email, subject, errVerify := auth.SupabaseTokenVerify(tokenIn)
	if errVerify != nil {
		return nil, huma.Error401Unauthorized("Invalid or expired token.")
	}

	userObj, errUser := userGetCreate(email, subject, ctx)
	if errUser != nil {
		return nil, errUser
	}

	tokenOut, errToken := auth.CreateJWT(strconv.Itoa(userObj.ID), userObj.Email)
	if errToken != nil {
		return nil, huma.Error400BadRequest("Unable to authenticate.")
	}

	response := &SignInSupabaseOutput{}
	response.Body.Token = tokenOut
	return response, nil
}

func userGetCreate(email string, subject string, ctx context.Context) (*ent.User, error) {
	if subject != "" {
		userObj, _ := db.EntDB.User.Query().
			Where(user.SupabaseEQ(subject)).
			Only(ctx)
		if userObj != nil {
			return userObj, nil
		}
	}

	userObj, _ := db.EntDB.User.Query().
		Where(user.EmailEqualFold(email)).
		Only(ctx)
	if userObj != nil {
		if subject != "" && userObj.Supabase != subject {
			_, _ = db.EntDB.User.
				UpdateOneID(userObj.ID).
				SetSupabase(subject).
				Save(ctx)
		}
		return userObj, nil
	}

	passwordPlain, errRandom := hexRandom(32)
	if errRandom != nil {
		return nil, huma.Error400BadRequest("Unable to create user.")
	}

	passwordHashed, errHash := password.HashPassword(passwordPlain)
	if errHash != nil {
		return nil, huma.Error400BadRequest("Unable to create user.")
	}

	userCreate := db.EntDB.User.Create().SetEmail(email).SetPassword(passwordHashed)
	if subject != "" {
		userCreate = userCreate.SetSupabase(subject)
	}
	userObj, errSave := userCreate.Save(ctx)
	if errSave != nil {
		return nil, huma.Error400BadRequest("Unable to create user.")
	}

	return userObj, nil
}

func hexRandom(byteCount int) (string, error) {
	bytesRandom := make([]byte, byteCount)
	if _, errRead := rand.Read(bytesRandom); errRead != nil {
		return "", errRead
	}
	return hex.EncodeToString(bytesRandom), nil
}

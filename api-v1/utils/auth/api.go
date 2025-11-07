package auth

import (
	"context"
	"errors"

	"api-go/ent"
	"api-go/ent/user"
	"api-go/utils/db"
)

type AuthParam struct {
	Auth string `header:"Authorization"`
	User *ent.User
}

func AuthUser(token string) (*ent.User, error) {
	ctx := context.Background()

	authID, authValid := GetJWT(token)
	if !authValid {
		return nil, errors.New("unable to authenticate")
	}

	userObj, err := db.EntDB.User.Query().Where(user.ID(authID)).Only(ctx)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return userObj, nil
}

package events

import (
	"context"
	"fmt"

	"api-go/ent/user"
	"api-go/utils/db"
)

func UserTrigger(userID int) {
	ctx := context.Background()
	userObj, err := db.EntDB.User.Query().
		Where(user.ID(userID)).
		Only(ctx)
	fmt.Println(userObj)
	fmt.Println(err)
}

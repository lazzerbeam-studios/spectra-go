package mutations

import (
	"context"
	"fmt"

	"api-go/ent"
	"api-go/ent/hook"
	"api-go/events"
	"api-go/utils/db"
)

func UserHook() {
	db.EntDB.User.Use(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, mutation *ent.UserMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, mutation)

			fmt.Println("User Hook")
			fmt.Println(mutation.Op())
			fmt.Println(mutation.ID())
			fmt.Println(mutation.IDs(ctx))
			fmt.Println(value)

			operationName := mutation.Op()
			userID, exists := mutation.ID()
			if exists && operationName.String() == "OpCreate" {
				go events.User_Created_Event(userID)
			}

			return value, err
		})
	})
}

---
name: huma-ent-api
description: Build CRUD REST APIs with Huma v2 and Ent—list, get one, update, delete—for global resources and for resources owned by the authenticated user (input.User). Use when adding or refactoring list/get/update/delete endpoints in Huma + Ent projects.
---

# Huma + Ent API

Patterns for **list**, **get one**, **update**, and **delete** with Huma v2 and Ent. Two variants: **global** (any caller) and **user-scoped** (resource must belong to **input.User**).

**Conventions**: Input/Output structs, **Resolve** for auth and validation, **huma.Error*** for errors. Replace **Item**/**item** and your auth/DB package names with your project’s names.

---

## Auth (when endpoints require a user)

Input embeds your auth param (e.g. **auth.AuthParam**) and resolves the user in **Resolve**:

```go
type YourGetInput struct {
	auth.AuthParam
	ID int `path:"id"`
}

func (input *YourGetInput) Resolve(ctx huma.Context) []error {
	var err error
	input.User, err = auth.AuthUser(input.Auth)
	if err != nil {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	return nil
}
```

Use **input.User** in the handler to scope queries to the current user.

---

## 1. List (get all) — global

* **Input**: **auth.AuthParam** (if needed), **query:"offset"**, **query:"limit"**, optional filters.
* **Output**: **Body.List** slice.
* **Query**: **db.EntDB.Item.Query().Where(...).Order().Offset().Limit().All(ctx)**.

```go
type ItemsGetInput struct {
	auth.AuthParam
	Offset int `query:"offset" default:"0"`
	Limit  int `query:"limit" default:"25" maximum:"100"`
}

type ItemsGetOutput struct {
	Body struct {
		List []models.Item `json:"list"`
	}
}

func ItemsGetAPI(ctx context.Context, input *ItemsGetInput) (*ItemsGetOutput, error) {
	items, err := db.EntDB.Item.Query().
		Where(item.Deleted(false)).
		Order(item.ByCreated(sql.OrderDesc())).
		Offset(input.Offset).
		Limit(input.Limit).
		All(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Items not found.")
	}
	response := &ItemsGetOutput{}
	list := &[]models.Item{}
	copier.Copy(&list, &items)
	response.Body.List = *list
	return response, nil
}
```

**Registration**: **GET /items** → **ItemsGetAPI**.

**Examples in repo**: **groups_api/op-GroupsGetAPI.go**, **questions_api/op-QuestionsGetAPI.go**.

---

## 2. List (get all) — user-scoped

* Same input/output shape; **query starts from the user** so only their resources are returned.
* Use the relation from User to your entity (e.g. **input.User.QueryItems()** or **QueryX().QueryItems()**).

```go
func UserItemsGetAPI(ctx context.Context, input *UserItemsGetInput) (*UserItemsGetOutput, error) {
	items, err := input.User.QueryItems().
		Where(item.Deleted(false)).
		Order(item.ByCreated(sql.OrderDesc())).
		Offset(input.Offset).
		Limit(input.Limit).
		All(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Items not found.")
	}
	response := &UserItemsGetOutput{}
	list := &[]models.Item{}
	copier.Copy(&list, &items)
	response.Body.List = *list
	return response, nil
}
```

**Examples in repo**: **educator_groups_api/op-EducatorGroupsGetAPI.go** (via **input.User.QueryEducatorClassrooms().QueryGroup()**), **users_api/op-ProfileNotificationsGetAPI.go**.

---

## 3. Get one — global

* **Input**: **auth.AuthParam** (if needed), **path:"id"**.
* **Output**: **Body.Object** single model.
* **Query**: **db.EntDB.Item.Query().Where(item.ID(input.ID)).With(...).Only(ctx)**.

```go
type ItemGetInput struct {
	auth.AuthParam
	ID int `path:"id"`
}

type ItemGetOutput struct {
	Body struct {
		Object models.Item `json:"object"`
	}
}

func ItemGetAPI(ctx context.Context, input *ItemGetInput) (*ItemGetOutput, error) {
	itemObj, err := db.EntDB.Item.Query().
		Where(item.Deleted(false)).
		Where(item.ID(input.ID)).
		WithRelation().
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Item not found.")
	}
	response := &ItemGetOutput{}
	object := &models.Item{}
	copier.Copy(&object, &itemObj)
	response.Body.Object = *object
	return response, nil
}
```

**Registration**: **GET /items/{id}** → **ItemGetAPI**.

**Examples in repo**: **groups_api/op-GroupGetAPI.go**, **assignments_api/op-AssignmentGetAPI.go**, **answers_api/op-AnswerGetAPI.go**.

---

## 4. Get one — user-scoped

* **Load only if the item belongs to the user**: start from **input.User.QueryItems()** (or the appropriate user→entity path).

```go
func UserItemGetAPI(ctx context.Context, input *UserItemGetInput) (*UserItemGetOutput, error) {
	itemObj, err := input.User.QueryItems().
		Where(item.Deleted(false)).
		Where(item.ID(input.ID)).
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Item not found.")
	}
	response := &UserItemGetOutput{}
	object := &models.Item{}
	copier.Copy(&object, &itemObj)
	response.Body.Object = *object
	return response, nil
}
```

**Examples in repo**: **educator_classrooms_api** (update/delete use **input.User.QueryEducatorClassrooms()**), **notifications_api/op-NotificationUpdateAPI.go** (**input.User.QueryReceiverNotifications()**).

---

## 5. Update — user-scoped (typical)

* **Input**: **auth.AuthParam**, **Body** with ID and updatable fields (e.g. from a shared **models.ItemUpdate**).
* **Steps**: (1) Load entity via user relation so only owner can update. (2) Apply **Update().SetX(...).Save(ctx)**. (3) Optionally re-query with **With(...)** and return full object.

```go
type ItemUpdateInput struct {
	auth.AuthParam
	Body struct {
		models.ItemUpdate
	}
}

type ItemUpdateOutput struct {
	Body struct {
		Object models.Item `json:"object"`
	}
}

func ItemUpdateAPI(ctx context.Context, input *ItemUpdateInput) (*ItemUpdateOutput, error) {
	itemObj, err := input.User.QueryItems().
		Where(item.Deleted(false)).
		Where(item.ID(input.Body.ID)).
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Item not found.")
	}
	itemObj, err = itemObj.Update().
		SetName(*input.Body.Name).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to update item.")
	}
	itemObj, err = input.User.QueryItems().
		Where(item.ID(input.Body.ID)).
		WithRelation().
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Item not found.")
	}
	response := &ItemUpdateOutput{}
	object := &models.Item{}
	copier.Copy(&object, &itemObj)
	response.Body.Object = *object
	return response, nil
}
```

**Registration**: **PUT /items** (or **PATCH**) with body containing ID.

**Examples in repo**: **educator_classrooms_api/op-EducatorClassroomUpdateAPI.go**, **notifications_api/op-NotificationUpdateAPI.go**, **users_api/op-ProfileUpdateAPI.go**.

---

## 6. Delete — user-scoped (typical)

* **Input**: **auth.AuthParam**, **path:"id"**.
* **Steps**: (1) Load entity via user relation. (2) Soft delete: **Update().SetDeleted(true).Save(ctx)** or hard delete: **Delete()**.

```go
type ItemDeleteInput struct {
	auth.AuthParam
	ID int `path:"id"`
}

type ItemDeleteOutput struct {
	Body struct {
		Object models.Item `json:"object"`
	}
}

func ItemDeleteAPI(ctx context.Context, input *ItemDeleteInput) (*ItemDeleteOutput, error) {
	itemObj, err := input.User.QueryItems().
		Where(item.Deleted(false)).
		Where(item.ID(input.ID)).
		Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Item not found.")
	}
	itemObj, err = itemObj.Update().
		SetDeleted(true).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to delete item.")
	}
	response := &ItemDeleteOutput{}
	object := &models.Item{}
	copier.Copy(&object, &itemObj)
	response.Body.Object = *object
	return response, nil
}
```

**Registration**: **DELETE /items/{id}** → **ItemDeleteAPI**.

**Examples in repo**: **educator_classrooms_api/op-EducatorClassroomDeleteAPI.go**, **answers_api/op-AnswerDeleteAPI.go**, **comments_api/op-CommentDeleteAPI.go**.

---

## 7. Registration (routes.go)

```go
func Register(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "ItemsGetAPI",
		Method:      http.MethodGet,
		Path:        "/items",
		Tags:        []string{"items"},
	}, ItemsGetAPI)

	huma.Register(api, huma.Operation{
		OperationID: "ItemGetAPI",
		Method:      http.MethodGet,
		Path:        "/items/{id}",
		Tags:        []string{"items"},
	}, ItemGetAPI)

	huma.Register(api, huma.Operation{
		OperationID: "ItemUpdateAPI",
		Method:      http.MethodPut,
		Path:        "/items",
		Tags:        []string{"items"},
	}, ItemUpdateAPI)

	huma.Register(api, huma.Operation{
		OperationID: "ItemDeleteAPI",
		Method:      http.MethodDelete,
		Path:        "/items/{id}",
		Tags:        []string{"items"},
	}, ItemDeleteAPI)
}
```

---

## Quick reference

| Operation   | Global query              | User-scoped query                    |
|------------|---------------------------|--------------------------------------|
| List       | **db.EntDB.Item.Query()...** | **input.User.QueryItems()...**         |
| Get one    | **db.EntDB.Item.Query().Where(item.ID(id)).Only(ctx)** | **input.User.QueryItems().Where(item.ID(id)).Only(ctx)** |
| Update     | —                         | Load via **input.User.QueryItems()**, then **obj.Update().SetX().Save(ctx)** |
| Delete     | —                         | Load via **input.User.QueryItems()**, then **obj.Update().SetDeleted(true).Save(ctx)** |

* Use **Resolve** to set **input.User** and return auth errors.
* Return **huma.Error404NotFound** when the entity is missing (or not owned); **huma.Error400BadRequest** for validation/update failures.
* Use **copier.Copy** (or your DTO mapper) from Ent types to response models.
* For user-scoped resources, **always** resolve the entity through a user relation (e.g. **input.User.QueryItems()**) so callers cannot access or modify other users’ data.

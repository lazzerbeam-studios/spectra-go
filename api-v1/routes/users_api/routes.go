package users_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "ProfileGetAPI",
		Method:      http.MethodGet,
		Path:        "/users/profile",
		Tags:        []string{"users"},
	}, ProfileGetAPI)

	huma.Register(api, huma.Operation{
		OperationID: "ProfileUpdateAPI",
		Method:      http.MethodPut,
		Path:        "/users/profile",
		Tags:        []string{"users"},
	}, ProfileUpdateAPI)

	huma.Register(api, huma.Operation{
		OperationID: "ProfileUpdateEmailAPI",
		Method:      http.MethodPut,
		Path:        "/users/profile/email",
		Tags:        []string{"users"},
	}, ProfileUpdateEmailAPI)

	huma.Register(api, huma.Operation{
		OperationID: "ProfileImageUpdateAPI",
		Method:      http.MethodPost,
		Path:        "/users/profile/image",
		Tags:        []string{"users"},
	}, ProfileImageUpdateAPI)

}

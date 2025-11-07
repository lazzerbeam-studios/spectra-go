package auth_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "SignUpAPI",
		Method:      http.MethodPost,
		Path:        "/auth/signup",
		Tags:        []string{"auth"},
	}, SignUpAPI)

	huma.Register(api, huma.Operation{
		OperationID: "SignInAPI",
		Method:      http.MethodPost,
		Path:        "/auth/signin",
		Tags:        []string{"auth"},
	}, SignInAPI)

	huma.Register(api, huma.Operation{
		OperationID: "PasswordForgotPostAPI",
		Method:      http.MethodPost,
		Path:        "/auth/passwordforgot",
		Tags:        []string{"auth"},
	}, PasswordForgotPostAPI)

	huma.Register(api, huma.Operation{
		OperationID: "PasswordResetPostAPI",
		Method:      http.MethodPost,
		Path:        "/auth/passwordreset",
		Tags:        []string{"auth"},
	}, PasswordResetPostAPI)

}

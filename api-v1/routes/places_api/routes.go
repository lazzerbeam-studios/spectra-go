package places_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "PlacesGetAPI",
		Method:      http.MethodGet,
		Path:        "/places",
		Tags:        []string{"places"},
	}, PlacesGetAPI)

	huma.Register(api, huma.Operation{
		OperationID: "PlaceGetAPI",
		Method:      http.MethodGet,
		Path:        "/places/{place_id}",
		Tags:        []string{"places"},
	}, PlaceGetAPI)

}

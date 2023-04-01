package routes

import (
	"api/src/controllers"
	"net/http"
)

var routePublicatons = []Route{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Func:           controllers.CreatePublication,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications",
		Method:         http.MethodGet,
		Func:           controllers.GetPublications,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodGet,
		Func:           controllers.GetPublication,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodPut,
		Func:           controllers.UpdatePublication,
		AuthIsRequired: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodDelete,
		Func:           controllers.DeletePublication,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}/publications",
		Method:         http.MethodGet,
		Func:           controllers.GetUserPublications,
		AuthIsRequired: true,
	},
}

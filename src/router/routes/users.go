package routes

import (
	"api/src/controllers"
	"net/http"
)

var routeUsers = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Func:           controllers.CreateUser,
		AuthIsRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Func:           controllers.GetUsers,
		AuthIsRequired: false,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Func:           controllers.GetUser,
		AuthIsRequired: false,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		AuthIsRequired: false,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		AuthIsRequired: false,
	},
}

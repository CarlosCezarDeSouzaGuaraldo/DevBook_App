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
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Func:           controllers.GetUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		AuthIsRequired: true,
	},
	{
		URI:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Func:           controllers.FollowUser,
		AuthIsRequired: true,
	},
}

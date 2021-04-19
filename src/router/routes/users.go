package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequestAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.FindAllUsers,
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequestAuth: false,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequestAuth: false,
	},
}

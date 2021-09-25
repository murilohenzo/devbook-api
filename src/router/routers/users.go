package routers

import (
	"api_resources/src/controllers"
	"net/http"
)

var usersRouters = []Router{
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.Create,
		Authentication: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.FindAll,
		Authentication: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.FindById,
		Authentication: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.Update,
		Authentication: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.Delete,
		Authentication: false,
	},
}

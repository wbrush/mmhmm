package api

import (
	"github.com/urfave/negroni"
	"github.com/wbrush/mmhmm/configuration"
)

const (
	NotesPath = "/notes"
	UserPath  = "/users"
)

func (api *API) initRoutes(wrapper *negroni.Negroni) {
	api.HandleActions(wrapper, configuration.APIBasePath, []Route{
		{
			Name:        "Info",
			Method:      "GET",
			Pattern:     "/info",
			HandlerFunc: api.HandleInfo,
			Middleware:  nil,
		},
		{
			Name:        "Ping",
			Method:      "GET",
			Pattern:     "/ping",
			HandlerFunc: api.HandlePing,
			Middleware:  nil,
		},
	})
	api.HandleActions(wrapper, configuration.APIBasePath+configuration.APIVersion, []Route{
		//  application specific - Users
		{
			Name:        "Create User",
			Method:      "POST",
			Pattern:     UserPath,
			HandlerFunc: api.CreateUser,
			Middleware:  nil,
			// Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		},
		{
			Name:        "Get User",
			Method:      "GET",
			Pattern:     UserPath + "/{id}",
			HandlerFunc: api.GetUser,
			Middleware:  nil,
			// Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		},
		{
			Name:        "List Users",
			Method:      "GET",
			Pattern:     UserPath,
			HandlerFunc: api.ListUsers,
			Middleware:  nil,
			// Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		},
		{
			Name:        "Update User",
			Method:      "PUT",
			Pattern:     UserPath + "/{id}",
			HandlerFunc: api.UpdateUser,
			Middleware:  nil,
			// Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		},
		{
			Name:        "Remove User",
			Method:      "DELETE",
			Pattern:     UserPath + "/{id}",
			HandlerFunc: api.DeleteUser,
			Middleware:  nil,
			// Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		},
	})
}

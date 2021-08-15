package api

import (
	"github.com/urfave/negroni"
	"github.com/wbrush/go-common/httphelper"
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
		//  application specific
		{
			Name:        "Create User",
			Method:      "POST",
			Pattern:     UserPath,
			HandlerFunc: api.CreateUser,
			Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		},
		// {
		// 	Name:        "Get Template",
		// 	Method:      "GET",
		// 	Pattern:     TemplatePath + "/{id}",
		// 	HandlerFunc: api.GetTemplate,
		// 	Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		// },
		// {
		// 	Name:        "List Templates",
		// 	Method:      "GET",
		// 	Pattern:     TemplatePath,
		// 	HandlerFunc: api.ListTemplates,
		// 	Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		// },
		// {
		// 	Name:        "Update Template",
		// 	Method:      "PUT",
		// 	Pattern:     TemplatePath + "/{id}",
		// 	HandlerFunc: api.UpdateTemplate,
		// 	Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		// },
		// {
		// 	Name:        "Remove Template",
		// 	Method:      "DELETE",
		// 	Pattern:     TemplatePath + "/{id}",
		// 	HandlerFunc: api.DeleteTemplate,
		// 	Middleware:  []negroni.HandlerFunc{httphelper.MWUserInfoHeader},
		// },
	})
}

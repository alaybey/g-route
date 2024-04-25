/*
The BaseController interface defines methods for handling common HTTP request
types such as GET, POST, PUT, DELETE, etc. This interface serves as a contract
for implementing controllers responsible for handling specific routes in the application.

The purpose of the BaseController interface is to provide a standardized way
of defining controllers and their corresponding methods, enabling consistency
and modularity in the routing approach of the application.

Usage:
    Implement this interface in your controller structs to define methods
    for handling various HTTP request types.

Example:
    type MyController struct{}

   	func NewMyController() *MyController {
		return &MyController{}
	}

	func (m *MyController) RegisterRoute(mux *http.ServeMux) {
		mux.HandleFunc("/myendpoint/", m.MyEndpointFunc)
	}
*/

package controller

import "net/http"

type BaseController interface {
	RegisterRoute(mux *http.ServeMux)
}

var registeredControllers []BaseController

func RegisterController(controller BaseController) {
	registeredControllers = append(registeredControllers, controller)
}

func GetRegisteredControllers() []BaseController {
	return registeredControllers
}

func RegisterRoutes(mux *http.ServeMux) {
	registeredControllers := GetRegisteredControllers()
	for _, controller := range registeredControllers {
		controller.RegisterRoute(mux)
	}
}

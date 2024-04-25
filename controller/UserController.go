package controller

import (
	"fmt"
	"net/http"
)

// Example
type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}
func (u *UserController) RegisterRoute(mux *http.ServeMux) {
	mux.HandleFunc("/user/", u.GetUser)
}
func (u *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Called GetUser")
}

func init() {
	RegisterController(&UserController{})
}

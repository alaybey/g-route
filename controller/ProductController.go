package controller

import (
	"fmt"
	"net/http"
)

// Example
type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (p *ProductController) RegisterRoute(mux *http.ServeMux) {
	mux.HandleFunc("/product/", p.GetProduct)
}

func (p *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Called GetProduct")
}

func init() {
	RegisterController(&ProductController{})
}

package api

import (
	"fmt"
	"net/http"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type API struct {
	usersRepo       repository.UserRepository
	productsRepo    repository.ProductRepository
	cartItemRepo    repository.CartItemRepository
	transactionRepo repository.TransactionRepository
	mux             *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, productsRepo, cartItemRepo, transactionRepo, mux,
	}
<<<<<<< HEAD
	// handle tiap rooting, root routing tidak ada

	mux.HandleFunc("/api/user/login", api.login)
	mux.HandleFunc("/api/user/logout", api.logout)
	mux.HandleFunc("/api/dashboard", api.dashboard)
	mux.HandleFunc("/api/products", api.productList)
	mux.HandleFunc("/api/cart/add", api.addToCart)
	mux.HandleFunc("/api/cart/clear", api.clearCart)
	mux.HandleFunc("/api/carts", api.cartList)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
=======

	mux.HandleFunc("/api/user/login", api.login)
	mux.HandleFunc("/api/user/logout", api.logout)

	mux.Handle("/api/dashboard", api.AuthMiddleWare(http.HandlerFunc(api.dashboard)))
	mux.Handle("/api/products", api.AuthMiddleWare(http.HandlerFunc(api.productList)))
	mux.Handle("/api/cart/add", api.AuthMiddleWare(http.HandlerFunc(api.addToCart)))

	// TODO: answer here

	// mux.HandleFunc("/api/dashboard", api.dashboard)
	// mux.HandleFunc("/api/products", api.productList)
	// mux.HandleFunc("/api/cart/add", api.addToCart)
	// mux.HandleFunc("/api/cart/clear", api.clearCart)

	// TODO: answer here
>>>>>>> c0397392214e368e84db7e7b9a1534ca43781bfb

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}

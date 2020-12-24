package samples

import (
	"net/http"

	"github.com/gorilla/mux"
)

func test1() {
	router := mux.NewRouter()
	router.Handle("/deploy", middleware.Auth(http.HandlerFunc(wa.serveDeploy), "deploy")).Methods("POST")
	router.Handle("/admin", http.HandlerFunc(wa.serveDeploy)).Methods("POST")
	router.PathPrefix("/auth/").HandlerFunc(wa.wrapAuth)
	router.PathPrefix("/login/").HandlerFunc(wa.wrapAuth)
	router.PathPrefix("/sensitive/").HandlerFunc(wa.wrapAuth)
}
func test2() {
	r := mux.NewRouter()
	r.Use(AuthMiddleware)
	r.Handle("/deploy", http.HandlerFunc(serveDeploy)).Methods("POST")
	r.PathPrefix("/auth/").HandlerFunc(wrapAuth)
}

func AuthMiddleware() {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

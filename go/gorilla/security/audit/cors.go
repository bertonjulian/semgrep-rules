package util

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func test() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", ac.SignUp).Methods("POST")
	router.HandleFunc("/signin", ac.SignIn).Methods("POST")
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk = handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func test1() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func test2() {
	r := mux.NewRouter()

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":8080", r)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	write.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}

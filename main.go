package main

import (
	"crud/dbconfig"
	"crud/routehandler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	options           string = "OPTIONS"
	allow_origin      string = "Access-Control-Allow-Origin"
	allow_methods     string = "Access-Control-Allow-Methods"
	allow_headers     string = "Access-Control-Allow-Headers"
	allow_credentials string = "Access-Control-Allow-Credentials"
	expose_headers    string = "Access-Control-Expose-Headers"
	credentials       string = "true"
	origin            string = "Origin"
	methods           string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"

	// If you want to expose some other headers add it here
	headers string = "Access-Control-Allow-Origin, Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token,ApiKey"
)

// Handler will allow cross-origin HTTP requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set allow origin to match origin of our request or fall back to *
		if o := r.Header.Get(origin); o != "" {
			w.Header().Set(allow_origin, o)
		} else {
			w.Header().Set(allow_origin, "*")
		}

		// Set other headers
		w.Header().Set(allow_headers, headers)
		w.Header().Set(allow_methods, methods)
		w.Header().Set(allow_credentials, credentials)
		w.Header().Set(expose_headers, headers)

		// If this was preflight options request let's write empty ok response and return
		if r.Method == options {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	a := mux.NewRouter()
	db, err := dbconfig.Connect()
	if err != nil {
		panic(err)
	}
	routehandler.HandleEmpRoutes(a, db)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", dbconfig.GetEnvironmentVars("PORT")), CORS(a)))

}

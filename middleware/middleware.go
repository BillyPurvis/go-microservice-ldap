package middleware

import (
	"net/http"

	"github.com/BillyPurvis/go-microservice-ldap/authenticate"
	"github.com/julienschmidt/httprouter"
)

// SetJSONHeader Intercept Responses
func SetJSONHeader(handler http.Handler) http.Handler {
	middlewareFun := func(w http.ResponseWriter, r *http.Request) {
		// Set type as JSON as we're no scrubs
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(middlewareFun)
}

// AuthenticateWare Protect Routes that require API Token
func AuthenticateWare(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		token := r.Header.Get("X-Api-Key")

		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		// Check Token
		authCheck := authenticate.TokenCheck(token)

		if authCheck != true {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			next(w, r, ps)
		}
	}
}

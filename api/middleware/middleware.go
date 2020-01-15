package middleware

import (
	"log"
	"net/http"
)

var (
	APIKey = "7d45d757a974e302bc54"
)

func ApplyMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

func LoggingMiddleware(next http.Handler) http.Handler {

	log.Println("logging middleware registered")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Here is the request URI:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {

	log.Println("logging middleware registered")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Here is the request URI:", r.Header.Get("X-Secret"))
		secret := r.Header.Get("X-Secret")

		if secret == APIKey {
			next.ServeHTTP(w, r)
		} else {
	 		http.Error(w, "Forbidden", http.StatusUnauthorized)
		}
	})
}

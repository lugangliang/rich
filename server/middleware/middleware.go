package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s, log", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

func Validate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s, valid", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

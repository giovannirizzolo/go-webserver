package logging

import (
	"log"
	"net/http"
)

var LoggingMiddleware http.Handler

func handleLogging(next http.Handler) {
	LoggingMiddleware = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println(r.RequestURI)

		next.ServeHTTP(w,r)
	})
}
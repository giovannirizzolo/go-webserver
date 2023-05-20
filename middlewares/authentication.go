package middlewares

import (
	"log"
	"net/http"
)

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

func Populate(amw *authenticationMiddleware) {
	amw.tokenUsers["e34242sad"] = "user0"
	amw.tokenUsers["dsad32rd"] = "user1"
	amw.tokenUsers["34f3fdffdfgv"] = "user2"
	amw.tokenUsers["fdsfdsfwe"] = "user3"
}

func (awm *authenticationMiddleware) AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("X-Session-Token")

		if user, found := awm.tokenUsers[token]; found {
			log.Printf("User authenticated %s", user)

			next.ServeHTTP(w,r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
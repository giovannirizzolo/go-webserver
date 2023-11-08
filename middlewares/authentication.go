package middlewares

import (
	"log"
	"net/http"
)

type AuthenticationMiddleware struct {
	TokenUsers map[string]string
}

func (amw *AuthenticationMiddleware)Populate() {
	amw.TokenUsers["e34242sad"] = "user0"
	amw.TokenUsers["dsad32rd"] = "user1"
	amw.TokenUsers["34f3fdffdfgv"] = "user2"
	amw.TokenUsers["fdsfdsfwe"] = "user3"
}

func (awm *AuthenticationMiddleware) AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		token := r.Header.Get("X-Session-Token")

		if user, found := awm.TokenUsers[token]; found {
			log.Printf("User authenticated %s", user)

			next.ServeHTTP(w,r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
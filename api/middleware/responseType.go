package middleware

import "net/http"

func SetHeaderContentType(contentType string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hf := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hf)
	}
}
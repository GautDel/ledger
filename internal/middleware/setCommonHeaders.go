package middleware

import "net/http"

func SetCommonHeaders(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headers := map[string]string{
			"Content-Type":  "application/json",
			"Cache-Control": "no-cache",
		}

		for key, value := range headers {
			w.Header().Set(key, value)
		}

        next.ServeHTTP(w, r)
	})
}


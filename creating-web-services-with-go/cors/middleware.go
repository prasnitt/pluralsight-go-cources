package cors

import "net/http"

// Middleware will be called before actual http handler request
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// perform following operations before calling actual handler
		w.Header().Add("Content-Type", "application/json")

		// enable CORS
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// call actual handler
		h.ServeHTTP(w, r)
	})
}

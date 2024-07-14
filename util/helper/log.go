package helper

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func URLRewriter(router *mux.Router, baseURLPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = func(url string) string {
			if strings.Index(url, baseURLPath) == 0 {
				url = url[len(baseURLPath):]
			}
			return url
		}(r.URL.Path)

		router.ServeHTTP(w, r)
	}
}

func LoggerMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "notifications") {
				next.ServeHTTP(w, r)
				return
			}

			start := time.Now()

			recorder := httptest.NewRecorder()
			next.ServeHTTP(recorder, r)

			for k, v := range recorder.Header() {
				w.Header()[k] = v
			}
			w.WriteHeader(recorder.Code)
			recorder.Body.WriteTo(w)

			responseTime := time.Since(start).Seconds()
			formattedResponseTime := fmt.Sprintf("%.9f", responseTime)
			formattedResponseTime = fmt.Sprintf("%sµs", formattedResponseTime)

			log.Printf("%s - [%s] - [%s] \"%s %s %s\" %d %s\n",
				r.RemoteAddr,
				time.Now().Format(time.RFC1123),
				formattedResponseTime,
				r.Method,
				r.URL.Path,
				r.Proto,
				recorder.Code,
				r.UserAgent(),
			)
		})
	}
}

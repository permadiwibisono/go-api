package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseObserver struct {
	http.ResponseWriter
	Request *http.Request
	code    int
	set     bool
}

func (w *ResponseObserver) Write(p []byte) (int, error) {
	if !w.set {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(p)
}

func (w *ResponseObserver) WriteHeader(code int) {
	w.set = true
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

var Logger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := &ResponseObserver{
			ResponseWriter: w, Request: r,
		}
		start := time.Now()
		defer func() {
			fmt.Printf("%s - %s %s %d (%v)\n", time.Now().Format("2006/02/01 15:04:05 Z07:00"), r.Method, r.URL.Path, o.code, time.Now().Sub(start))
		}()
		next.ServeHTTP(o, r)
	})
}

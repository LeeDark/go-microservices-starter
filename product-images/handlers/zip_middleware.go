package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// create a gzipped response
			wrw := NewWrappedResponseWriter(rw)
			rw.Header().Set("Content-Encoding", "gzip")
			defer wrw.Flush()

			// call the next handler with the wrapped response writer
			next.ServeHTTP(wrw, r)
			return
		}

		// handle normal
		next.ServeHTTP(rw, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw: rw, gw: gw}
}

func (w *WrappedResponseWriter) Header() http.Header {
	return w.rw.Header()
}

func (w *WrappedResponseWriter) Write(b []byte) (int, error) {
	return w.gw.Write(b)
}

func (w *WrappedResponseWriter) WriteHeader(statusCode int) {
	w.rw.WriteHeader(statusCode)
}

func (w *WrappedResponseWriter) Flush() {
	w.gw.Flush()
	w.gw.Close()
}

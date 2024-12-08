package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	// добавляем middleware для перенаправления
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Если путь начинается с /api/, возвращаем кастомный ответ
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello from API"))
			return
		}

		// Перенаправляем на Hugo
		
		link := fmt.Sprintf("http://%s:%s", rp.host, rp.port)
		uri, _ := url.Parse(link)
		proxy := httputil.NewSingleHostReverseProxy(uri)
		proxy.ServeHTTP(w, r)
		
	})
}

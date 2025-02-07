package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	// "path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// compile vuejs application in the web/ directory
//go:generate npm run build
//go:embed dist/*
var frontendFS embed.FS

func main() {

	const ollamaURLPrefix = "/ollama"
	// Target server URL
	targetServerURL, err := url.Parse("http://localhost:11434") // Replace with your upstream server
	if err != nil {
		log.Fatal("Error parsing target URL:", err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetServerURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetServerURL.Scheme
		req.URL.Host = targetServerURL.Host
		req.URL.Path = targetServerURL.Path + strings.TrimPrefix(req.URL.Path, ollamaURLPrefix)
	}
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("Error proxying request: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))
	}


	r := chi.NewRouter()
	r.Use(middleware.Logger)
	

	// ollama/ need to be routed to http://localhost:11434
	r.HandleFunc(ollamaURLPrefix+"/*", func(w http.ResponseWriter, r *http.Request) { // Catch-all route
		proxy.ServeHTTP(w, r)
	})

	// serves a directory outside the binary
	// fs := http.FileServer(http.Dir(filepath.Join("web", "dist")))
    // r.Handle("/*", http.StripPrefix("/", fs))

	// we need to strip off the directory prefix from the embed.FS with fs.Sub
	embededFiles, _ := fs.Sub(frontendFS, "dist")
	// r.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(embededFiles))))
	r.Handle("/*", http.FileServer(http.FS(embededFiles)))
	
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome this is a test"))
	// })


	http.ListenAndServe(":3000", r)
}

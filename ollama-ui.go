package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	// "path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// compile vuejs application in the web/ directory
//go:generate npm run build
//go:embed dist/*
var frontendFS embed.FS

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
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
	fmt.Print("Listening on port http://localhost:3000")
}

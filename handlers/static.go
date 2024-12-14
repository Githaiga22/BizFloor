package handlers

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleErrors(w, http.StatusMethodNotAllowed)
		return
	}
	// Remove the /static/ prefix from the URL path
	filePath := path.Join("static", strings.TrimPrefix(r.URL.Path, "/static/"))
	// Check if the file exists and is not a directory
	info, err := os.Stat(filePath)
	if err != nil || info.IsDir() {
		HandleErrors(w, http.StatusNotFound)
		return
	}
	// Check the file extension
	ext := filepath.Ext(filePath)
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".otf":
		w.Header().Set("Content-Type", "font/otf")
	default:
		HandleErrors(w, http.StatusNotFound)
		return
	}
	// Serve the file
	http.ServeFile(w, r, filePath)
}

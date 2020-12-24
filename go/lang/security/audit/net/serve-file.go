package util

import (
	"net/http"
	"path/filepath"
)

func test() {
	var w http.ResponseWriter
	var r *http.Request
	dir := "./build/static/"
	file := "user input"
	http.ServeFile(w, r, filepath.Join(dir, "index.html"))
	http.ServeFile(w, r, filepath.Join(dir, file))
	http.ServeFile(w, r, file)
	http.ServeFile(w, r, file+dir)
}

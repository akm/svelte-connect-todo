package images

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
)

//go:embed files/*
var files embed.FS

func GetImage(w http.ResponseWriter, r *http.Request) {
	filename := fmt.Sprintf("files/number-%s.png", r.PathValue("id"))
	reader, err := files.Open(filename)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer reader.Close()
	io.Copy(w, reader)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filename))
	w.WriteHeader(http.StatusOK)
}

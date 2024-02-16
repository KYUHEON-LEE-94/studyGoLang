package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../wwwroot")))
	//http.Handle("/static", http.FileServer(http.Dir("wwwroot")))
	http.ListenAndServe(":5000", nil)
}

// func main() {
// 	http.Handle("/", new(staticHandler))

// 	http.ListenAndServe(":5000", nil)
// }

// type staticHandler struct {
// 	http.Handler
// }

// func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	print(req.URL.Path)
// 	localPath := "../wwwroot" + req.URL.Path
// 	content, err := os.ReadFile(localPath)
// 	if err != nil {
// 		w.WriteHeader(404)
// 		w.Write([]byte(http.StatusText(404)))
// 		return
// 	}

// 	contentType := getContentType(localPath)
// 	w.Header().Add("Content-Type", contentType)
// 	w.Write(content)
// }

// func getContentType(localPath string) string {
// 	var contentType string
// 	ext := filepath.Ext(localPath)

// 	switch ext {
// 	case ".html":
// 		contentType = "text/html"
// 	case ".css":
// 		contentType = "text/css"
// 	case ".js":
// 		contentType = "application/javascript"
// 	case ".png":
// 		contentType = "image/png"
// 	case ".jpg":
// 		contentType = "image/jpeg"
// 	default:
// 		contentType = "text/plain"
// 	}

// 	return contentType
// }

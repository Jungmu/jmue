package staticHandler

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/Jungmu/jmue/logger"
)

//Handler : http.Handler wrapping
type Handler struct {
	http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger.Debug("Get Request - URL : " + req.Host + req.URL.Path)
	logger.Debug("Get Request - User Info : " + req.UserAgent())

	localPath := ""

	if req.URL.Path == "/" {
		localPath = "wwwroot/index.html"
	} else {
		localPath = "wwwroot" + req.URL.Path
	}
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404) + "\n페이지를 찾을 수 없습니다."))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}

package staticHandler

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Jungmu/jmue/logger"
)

//Handler : http.Handler wrapping
type Handler struct {
	http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger.Debug("Get Request - URL : " + req.Host + req.URL.Path)
	logger.Debug("Get Request - User Info : " + req.UserAgent())

	localPath := route(req.URL.Path, w, req)

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

func route(path string, w http.ResponseWriter, req *http.Request) string {
	var routingPath string
	if path == "/" {
		routingPath = "wwwroot/index.html"
	} else {
		fileType := strings.Split(path, ".")

		if isStaticFile(fileType[len(fileType)-1]) {
			routingPath = "wwwroot" + path
		} else {
			routingPath = "wwwroot" + path + ".html"
		}
	}

	return routingPath
}

func isStaticFile(fileType string) bool {
	staticFileList := []string{"css", "js", "html"}
	staticFileList = append(staticFileList, "jpg", "jpeg", "png", "gif")

	for _, str := range staticFileList {
		if strings.EqualFold(fileType, str) {
			return true
		}
	}
	return false
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

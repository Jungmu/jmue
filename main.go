package main

import (
	"fmt"
	"jmue/logger"
	"jmue/staticHandler"
	"net/http"
	"os"
	"time"
)

func main() {
	now := time.Now()
	logFilePath := fmt.Sprintf("wwwroot/log/%d-%02d.log", now.Year(), now.Month())
	fpLog, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	logger.InitLogger(fpLog)

	http.Handle("/", new(staticHandler.Handler))
	http.Handle("/static", http.FileServer(http.Dir("wwwroot/static")))
	http.ListenAndServe(":80", nil)
}

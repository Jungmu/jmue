package main

import (
	"fmt"
	"jmue/jmueConst"
	"jmue/logger"
	"jmue/staticHandler"
	"net/http"
	"os"
	"time"
)

var (
	mode jmueConst.Mode
)

func main() {
	checkMode()

	now := time.Now()
	logFilePath := fmt.Sprintf("wwwroot/log/%d-%02d.log", now.Year(), now.Month())
	fpLog, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	logger.InitLogger(fpLog, mode)

	http.Handle("/", new(staticHandler.Handler))
	http.Handle("/static", http.FileServer(http.Dir("wwwroot/static")))
	http.ListenAndServe(":80", nil)
}

func checkMode() {
	if len(os.Args) < 2 {
		mode = jmueConst.SERVICE
	} else if os.Args[1] == "debug" || os.Args[1] == "d" {
		mode = jmueConst.DEBUG
	} else if os.Args[1] == "test" || os.Args[1] == "t" {
		mode = jmueConst.TEST
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Jungmu/jmue/pkg/const"
	"github.com/Jungmu/jmue/pkg/logger"
	"github.com/Jungmu/jmue/pkg/staticHandler"
)

var (
	mode jmueConst.Mode
)

// Home : for /
type Home struct {
	Title string
}

func newHome(title string) *Home {
	home := Home{}
	home.Title = title
	return &home
}

func main() {
	checkMode()

	now := time.Now()
	logFilePath := fmt.Sprintf("wwwroot/log/%d-%02d-%02d_%02d.log", now.Year(), now.Month(), now.Day(), now.Hour())
	fpLog, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	logger.InitLogger(fpLog, mode)

	http.Handle("/", staticHandler.New(newHome("Home")))
	http.Handle("/page", staticHandler.New(newHome("Page")))
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

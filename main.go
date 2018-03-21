package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/Jungmu/jmue/pkg/const"
	"github.com/Jungmu/jmue/pkg/logger"
	"github.com/Jungmu/jmue/pkg/staticHandler"
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
	http.HandleFunc("/gitPush", gitPushHandler)
	http.HandleFunc("/restart", serverRestart)
	http.Handle("/static", http.FileServer(http.Dir("wwwroot/static")))

	http.ListenAndServe(":80", nil)
}

func gitPushHandler(w http.ResponseWriter, req *http.Request) {
	execCommand("sh", "gitPush.sh")
	w.Write([]byte("OK"))
}

func serverRestart(w http.ResponseWriter, req *http.Request) {
	execCommand("sh", "gitPush.sh")
	w.Write([]byte("please wait, restart server..."))
}

func execCommand(command string, option string) {
	cmd := exec.Command(command, option)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
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

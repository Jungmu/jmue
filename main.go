package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
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
	logFilePath := fmt.Sprintf("wwwroot/log/%d-%02d-%02d.log", now.Year(), now.Month(), now.Day())
	fpLog, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()
	logger.InitLogger(fpLog, mode)

	http.HandleFunc("/sendEmail", sendEmail)

	http.Handle("/", new(staticHandler.Handler))
	http.Handle("/static", http.FileServer(http.Dir("wwwroot/static")))

	http.ListenAndServe(":80", nil)
}

func sendEmail(w http.ResponseWriter, req *http.Request) {
	pw, fileReadErr := ioutil.ReadFile("pw.jmue")
	if fileReadErr != nil {
		panic(fileReadErr)
	}
	log.Println(pw)
	auth := smtp.PlainAuth("", "whdrjs0@gmail.com", string(pw), "smtp.gmail.com")

	from := "admin.jmue.com"
	to := []string{"whdrjs0@gmail.com"}
	msg := []byte("To: whdrjs0@gmail.com\r\n" +
		"Subject: mail test!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, req, "http://jmue.xyz?send=ok", 301)
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

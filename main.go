package main

import (
	"fmt"
	"io/ioutil"
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

type emailStruct struct {
	to      []string
	from    string
	subject string
	msg     string
}

func main() {
	checkMode()

	now := time.Now()
	logFilePath := fmt.Sprintf("wwwroot/log/%d-%02d-%02d.log", now.Year(), now.Month(), now.Day())
	fpLog, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Error("log file open fail")
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
	req.ParseForm()

	var mailContents emailStruct
	parseMailContemts(&mailContents, w, req)

	pw := readAllfileText("pw.jmue")

	sendToStmp(pw, mailContents)

}

func parseMailContemts(mailContents *emailStruct, w http.ResponseWriter, req *http.Request) {
	mailContents.to = []string{"whdrjs0@gmail.com"}
	mailContents.from = req.FormValue("from")
	mailContents.subject = req.FormValue("subject")
	mailContents.msg = req.FormValue("msg")

	if mailContents.from == "" {
		http.Redirect(w, req, "http://jmue.xyz?send=fail", 301)
	}
}

func readAllfileText(fileName string) string {
	pw, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Error("file read fail" + err.Error())
		return ""
	}
	return string(pw)
}

func sendToStmp(pw string, mailContents emailStruct) {
	auth := smtp.PlainAuth("", mailContents.to[0], pw, "smtp.gmail.com")

	msg := []byte("To: " + mailContents.to[0] + "\r\n" +
		"Subject: " + mailContents.subject + "\r\n" +
		"\r\n" +
		"From: " + mailContents.from + "\r\n" +
		mailContents.msg + "\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, mailContents.from, mailContents.to, msg)
	if err != nil {
		logger.Error("send email fail" + err.Error())
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

package main

import (
	"crypto/tls"
	"fmt"
	"jmue/jmueConst"
	"jmue/logger"
	"jmue/staticHandler"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/acme/autocert"
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

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("jmue.xyz"), //Your domain here
		Cache:      autocert.DirCache("certs"),         //Folder for storing certificates
	}

	http.Handle("/", new(staticHandler.Handler))
	http.Handle("/static", http.FileServer(http.Dir("wwwroot/static")))

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	log.Fatal(server.ListenAndServeTLS("/etc/letsencrypt/live/jmue.xyz/fullchain.pem", "/etc/letsencrypt/live/jmue.xyz/privkey.pem")) //Key and cert are coming from Let's Encrypt
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

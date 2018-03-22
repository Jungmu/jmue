package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	pushCount int
)

func main() {
	pushCount = 0

	http.HandleFunc("/restart", serverRestart)
	http.HandleFunc("/gitPush", gitPushHandler)
	http.ListenAndServe(":3000", nil)
}

func serverRestart(w http.ResponseWriter, req *http.Request) {
	if pushCount > 0 {
		execCommand("restart.bat")
		w.Write([]byte("<h1>restart complete</h1><br>"))
	} else {
		w.Write([]byte("<h1>we don't need restart!</h1><br>"))
	}
	pushCount = 0
	w.Header().Add("Content-Type", "text/html")
	http.Redirect(w, req, "http://jmue.xyz", 301)
}

func gitPushHandler(w http.ResponseWriter, req *http.Request) {
	pushCount++
	execCommand("gitPull.bat")
	execCommand("sassBuild.bat")
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte("<h1>git pull complete</h1><br>"))
	http.Redirect(w, req, "http://jmue.xyz", 301)
}

func execCommand(command string) {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

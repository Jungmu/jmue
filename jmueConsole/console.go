package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

var (
	pushCount int
)

func main() {
	pushCount = 0

	go hotReloadSass()

	http.HandleFunc("/restart", serverRestart)
	http.HandleFunc("/gitPush", gitPushHandler)
	http.ListenAndServe(":3000", nil)
}

func hotReloadSass() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					execCommand("sassBuild.bat")
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("../scss")
	if err != nil {
		log.Fatal(err)
	}
	<-done
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

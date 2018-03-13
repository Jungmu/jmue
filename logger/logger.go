package logger

import (
	"io"
	"log"
	"os"
)

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
)

func init() {

}

//InitLogger : logging file and stdout
func InitLogger(fpLog *os.File) {
	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.LstdFlags)
	debugLogger.SetOutput(multiWriter)
	errorLogger.SetOutput(multiWriter)

	Debug("start server - initLogger complete")
}

//Debug : print debug text
func Debug(text string) {
	debugLogger.Printf(text)
}

//Error : print error text
func Error(text string) {
	errorLogger.Printf(text)
}

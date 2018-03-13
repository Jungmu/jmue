package logger

import (
	"io"
	"jmue/jmueConst"
	"log"
	"os"
)

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
	testLogger  *log.Logger
	loggingmode jmueConst.Mode
)

func init() {

}

//InitLogger : logging file and stdout
func InitLogger(fpLog *os.File, mode jmueConst.Mode) {
	loggingmode = mode
	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.LstdFlags)
	testLogger = log.New(os.Stdout, "TEST: ", log.LstdFlags)
	debugLogger.SetOutput(multiWriter)
	errorLogger.SetOutput(multiWriter)
	testLogger.SetOutput(os.Stdout)

	Debug("start server - initLogger complete")
}

//Debug : print debug text console and file
func Debug(text string) {
	if loggingmode != jmueConst.SERVICE {
		debugLogger.Printf(text)
	}
}

//Error : print error text console and file
func Error(text string) {
	errorLogger.Printf(text)
}

//Test : print error text just console
func Test(text string) {
	if loggingmode == jmueConst.TEST {
		testLogger.Printf(text)
	}
}

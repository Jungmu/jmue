package logger

import (
	"io"
	"log"
	"os"

	"github.com/Jungmu/jmue/pkg/const"
	"github.com/Jungmu/jmue/pkg/whereami"
)

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
	testLogger  *log.Logger
	loggingmode jmueConst.Mode
	isInit      bool
)

func init() {
	isInit = false
}

//InitLogger : logging file and stdout
func InitLogger(fpLog *os.File, mode jmueConst.Mode) {
	loggingmode = mode
	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	debugLogger = log.New(os.Stdout, "DEBUG: "+whereami.WhereAmI(), log.LstdFlags)
	errorLogger = log.New(os.Stdout, "ERROR: "+whereami.WhereAmIforError(), log.LstdFlags)
	testLogger = log.New(os.Stdout, "TEST: "+whereami.WhereAmI(), log.LstdFlags)
	debugLogger.SetOutput(multiWriter)
	errorLogger.SetOutput(multiWriter)
	testLogger.SetOutput(os.Stdout)

	isInit = true
	Debug("start server - initLogger complete")
}

func checkInitLogger() bool {
	if isInit == false {
		log.Println("need InitLogger(fpLog *os.File, mode jmueConst.Mode)")
		return false
	}
	return true
}

//Debug : print debug text console and file
func Debug(text string) {
	if checkInitLogger() == false {
		return
	}
	if loggingmode != jmueConst.SERVICE {
		debugLogger.Printf(text)
	}
}

//Error : print error text console and file
func Error(text string) {
	if checkInitLogger() == false {
		return
	}
	errorLogger.Printf(text)
}

//Test : print error text just console
func Test(text string) {
	if checkInitLogger() == false {
		return
	}
	if loggingmode == jmueConst.TEST {
		testLogger.Printf(text)
	}
}

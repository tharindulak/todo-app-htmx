package log

import (
	"log"
	"os"
)

var (
	EnableDebug = false
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func init() {
	infoLogger.SetOutput(os.Stdout)
	debugLogger.SetOutput(os.Stdout)
	errorLogger.SetOutput(os.Stderr)

	if envDebug, ok := os.LookupEnv("DEBUG"); ok {
		EnableDebug = envDebug == "true"
	}
}

func Info(v ...interface{}) {
    infoLogger.Println(v...)
}

func Debug(v ...interface{}) {
    if EnableDebug {
        debugLogger.Println(v...)
    }
}

func Error(v ...interface{}) {
    errorLogger.Println(v...)
}


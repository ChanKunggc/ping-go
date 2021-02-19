package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	DEBUGLevel = iota
	WARNLevel
	INFOLevel
	ERRORLevel
	DEBUG = "DEBUG"
	WARN  = "WARN"
	INFO  = "INFO"
	ERROR = "ERROR"
)

/*根据cmd 入参 -v 或--verbose而变化*/
var Verbose = false

func Info(msg interface{}) {

	logger.Printf("[%v] %v", INFO, msg)
}
func Infof(format string, v ...interface{}) {
	logger.Printf("[%v] %v", INFO, fmt.Sprintf(format, v...))
}
func Error(msg interface{}) {
	logger.Printf("[%v] %v", ERROR, msg)
}
func Errorf(format string, v ...interface{}) {

	logger.Printf("[%v] %v", ERROR, fmt.Sprintf(format, v...))
}
func Debug(msg interface{}) {
	if !Verbose {
		return
	}
	logger.Printf("[%v] %v", DEBUG, msg)
}
func Debugf(format string, v ...interface{}) {
	if !Verbose {
		return
	}
	logger.Printf("[%v] %v", DEBUG, fmt.Sprintf(format, v...))
}
func Warn(msg interface{}) {
	if !Verbose {
		return
	}
	logger.Printf("[%v] %v", WARN, msg)
}
func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

var logger *log.Logger

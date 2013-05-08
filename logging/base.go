package logging

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

type Logger interface {
	Error(v ...interface{}) string
	Errorf(format string, v ...interface{}) string
	Errorln(v ...interface{}) string
	Fatal(v ...interface{}) string
	Fatalf(format string, v ...interface{}) string
	Fatalln(v ...interface{}) string
	Info(v ...interface{}) string
	Infof(format string, v ...interface{}) string
	Infoln(v ...interface{}) string
	Panic(v ...interface{}) string
	Panicf(format string, v ...interface{}) string
	Panicln(v ...interface{}) string
	Warn(v ...interface{}) string
	Warnf(format string, v ...interface{}) string
	Warnln(v ...interface{}) string
}

func getInvokerLocation(skipNumber int) string {
	pc, file, line, ok := runtime.Caller(skipNumber)
	if !ok {
		return ""
	}
	simpleFileName := ""
	if index := strings.LastIndex(file, "/"); index > 0 {
		simpleFileName = file[index+1 : len(file)]
	}
	funcPath := ""
	funcPtr := runtime.FuncForPC(pc)
	if funcPtr != nil {
		funcPath = funcPtr.Name()
	}
	return fmt.Sprintf("%s : (%s:%d)", funcPath, simpleFileName, line)
}

func generateLogContent(
	logTag LogTag,
	invokingGapNumber uint,
	format string,
	v ...interface{}) string {
	skipNumber := int(invokingGapNumber) + 2
	baseInfo :=
		fmt.Sprintf("%s %s - ", logTag.Prefix(), getInvokerLocation(skipNumber))
	var result string
	if len(format) > 0 {
		result = fmt.Sprintf((baseInfo + format), v...)
	} else {
		vLen := len(v)
		params := make([]interface{}, (vLen + 1))
		params[0] = baseInfo
		for i := 1; i <= vLen; i++ {
			params[i] = v[i-1]
		}
		result = fmt.Sprint(params...)
	}
	return result
}

func GetSimpleLogger() Logger {
	return GetLogger([]Logger{&ConsoleLogger{invokingNumber: 2}})
}

func GetLogger(loggers []Logger) Logger {
	return &LogManager{loggers: loggers}
}

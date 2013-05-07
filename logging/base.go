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
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})
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
	callMidHierarchy uint,
	format string,
	v ...interface{}) string {
	skipNumber := int(callMidHierarchy) + 2
	baseInfo :=
		fmt.Sprintf("%s %s - ", logTag.Prefix(), getInvokerLocation(skipNumber))
	var result string
	if len(format) > 0 {
		result = fmt.Sprintf((baseInfo + format), v)
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
	return GetLogger([]Logger{&ConsoleLogger{}})
}

func GetLogger(loggers []Logger) Logger {
	return &LogManager{loggers: loggers}
}

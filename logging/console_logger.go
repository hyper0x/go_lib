package logging

import (
	"log"
)

type ConsoleLogger struct {
}

func (logger *ConsoleLogger) Error(v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), 1, "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Errorf(format string, v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), 1, format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Errorln(v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), 1, "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Fatal(v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), 1, "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Fatalf(format string, v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), 1, format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Fatalln(v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), 1, "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Info(v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), 1, "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Infof(format string, v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), 1, format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Infoln(v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), 1, "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Panic(v ...interface{}) {
	content := generateLogContent(getPanicLogTag(), 1, "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Panicf(format string, v ...interface{}) {
	content := generateLogContent(getPanicLogTag(), 1, format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Panicln(v ...interface{}) {
	content := generateLogContent(getPanicLogTag(), 1, "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Warn(v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), 1, "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Warnf(format string, v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), 1, format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Warnln(v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), 1, "", v)
	log.Println(content)
}

package logging

import (
	"log"
)

type ConsoleLogger struct {
	invokeNumber uint
}

func (logger *ConsoleLogger) SetInvokeNumber(invokeNumber uint) {
	logger.invokeNumber = invokeNumber
}

func (logger *ConsoleLogger) getInvokeNumber() uint {
	return logger.invokeNumber
}

func (logger *ConsoleLogger) SetDefaultInvokeNumber() {
	logger.invokeNumber = 1
}

func (logger *ConsoleLogger) Error(v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), logger.getInvokeNumber(), "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Errorf(format string, v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), logger.getInvokeNumber(), format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Errorln(v ...interface{}) {
	content := generateLogContent(getErrorLogTag(), logger.getInvokeNumber(), "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Fatal(v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), logger.getInvokeNumber(), "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Fatalf(format string, v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), logger.getInvokeNumber(), format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Fatalln(v ...interface{}) {
	content := generateLogContent(getFatalLogTag(), logger.getInvokeNumber(), "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Info(v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), logger.getInvokeNumber(), "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Infof(format string, v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), logger.getInvokeNumber(), format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Infoln(v ...interface{}) {
	content := generateLogContent(getInfoLogTag(), logger.getInvokeNumber(), "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Panic(v ...interface{}) {
	content := generateLogContent(getPanicLogTag(), logger.getInvokeNumber(), "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Panicf(format string, v ...interface{}) {
	content := generateLogContent(getPanicLogTag(), logger.getInvokeNumber(), format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Panicln(v ...interface{}) {
	content := generateLogContent(getPanicLogTag(), logger.getInvokeNumber(), "", v)
	log.Println(content)
}

func (logger *ConsoleLogger) Warn(v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), logger.getInvokeNumber(), "", v)
	log.Print(content)
}

func (logger *ConsoleLogger) Warnf(format string, v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), logger.getInvokeNumber(), format, v)
	log.Print(content)
}

func (logger *ConsoleLogger) Warnln(v ...interface{}) {
	content := generateLogContent(getWarnLogTag(), logger.getInvokeNumber(), "", v)
	log.Println(content)
}

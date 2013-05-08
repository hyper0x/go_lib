package logging

import (
	"log"
)

type ConsoleLogger struct {
	invokingNumber uint
}

func (logger *ConsoleLogger) getInvokingNumber() uint {
	return logger.invokingNumber
}

func (logger *ConsoleLogger) SetInvokingNumber(invokingNumber uint) {
	logger.invokingNumber = invokingNumber
}

func (logger *ConsoleLogger) SetDefaultInvokingNumber() {
	logger.invokingNumber = 1
}

func (logger *ConsoleLogger) Error(v ...interface{}) string {
	content := generateLogContent(getErrorLogTag(), logger.getInvokingNumber(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Errorf(format string, v ...interface{}) string {
	content := generateLogContent(getErrorLogTag(), logger.getInvokingNumber(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Errorln(v ...interface{}) string {
	content := generateLogContent(getErrorLogTag(), logger.getInvokingNumber(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Fatal(v ...interface{}) string {
	content := generateLogContent(getFatalLogTag(), logger.getInvokingNumber(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Fatalf(format string, v ...interface{}) string {
	content := generateLogContent(getFatalLogTag(), logger.getInvokingNumber(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Fatalln(v ...interface{}) string {
	content := generateLogContent(getFatalLogTag(), logger.getInvokingNumber(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Info(v ...interface{}) string {
	content := generateLogContent(getInfoLogTag(), logger.getInvokingNumber(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Infof(format string, v ...interface{}) string {
	content := generateLogContent(getInfoLogTag(), logger.getInvokingNumber(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Infoln(v ...interface{}) string {
	content := generateLogContent(getInfoLogTag(), logger.getInvokingNumber(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Panic(v ...interface{}) string {
	content := generateLogContent(getPanicLogTag(), logger.getInvokingNumber(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Panicf(format string, v ...interface{}) string {
	content := generateLogContent(getPanicLogTag(), logger.getInvokingNumber(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Panicln(v ...interface{}) string {
	content := generateLogContent(getPanicLogTag(), logger.getInvokingNumber(), "", v...)
	log.Println(content)
	return content
}

func (logger *ConsoleLogger) Warn(v ...interface{}) string {
	content := generateLogContent(getWarnLogTag(), logger.getInvokingNumber(), "", v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Warnf(format string, v ...interface{}) string {
	content := generateLogContent(getWarnLogTag(), logger.getInvokingNumber(), format, v...)
	log.Print(content)
	return content
}

func (logger *ConsoleLogger) Warnln(v ...interface{}) string {
	content := generateLogContent(getWarnLogTag(), logger.getInvokingNumber(), "", v...)
	log.Println(content)
	return content
}

package logging

type LogManager struct {
	loggers []Logger
}

func (self *LogManager) Error(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Error(v)
	}
}

func (self *LogManager) Errorf(format string, v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Errorf(format, v)
	}
}

func (self *LogManager) Errorln(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Errorln(v)
	}
}

func (self *LogManager) Fatal(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Fatal(v)
	}
}

func (self *LogManager) Fatalf(format string, v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Fatalf(format, v)
	}
}

func (self *LogManager) Fatalln(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Fatalln(v)
	}
}

func (self *LogManager) Info(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Info(v)
	}
}

func (self *LogManager) Infof(format string, v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Infof(format, v)
	}
}

func (self *LogManager) Infoln(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Infoln(v)
	}
}

func (self *LogManager) Panic(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Panic(v)
	}
}

func (self *LogManager) Panicf(format string, v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Panicf(format, v)
	}
}

func (self *LogManager) Panicln(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Panicln(v)
	}
}

func (self *LogManager) Warn(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Warn(v)
	}
}

func (self *LogManager) Warnf(format string, v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Warnf(format, v)
	}
}

func (self *LogManager) Warnln(v ...interface{}) {
	for _, logger := range self.loggers {
		logger.Warnln(v)
	}
}

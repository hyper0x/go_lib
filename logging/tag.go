package logging

const (
	ERROR_LOG_TAG = "ERROR"
	FATAL_LOG_TAG = "FATAL"
	INFO_LOG_TAG  = "INFO"
	PANIC_LOG_TAG = "PANIC"
	WARN_LOG_TAG  = "WARN"
)

type LogTag struct {
	name   string
	prefix string
}

func (self *LogTag) Name() string {
	return self.name
}

func (self *LogTag) Prefix() string {
	return self.prefix
}

var logTagMap map[string]LogTag = map[string]LogTag{
	ERROR_LOG_TAG: LogTag{name: ERROR_LOG_TAG, prefix: "[" + ERROR_LOG_TAG + "]"},
	FATAL_LOG_TAG: LogTag{name: FATAL_LOG_TAG, prefix: "[" + FATAL_LOG_TAG + "]"},
	INFO_LOG_TAG:  LogTag{name: INFO_LOG_TAG, prefix: "[" + INFO_LOG_TAG + "]"},
	PANIC_LOG_TAG: LogTag{name: PANIC_LOG_TAG, prefix: "[" + PANIC_LOG_TAG + "]"},
	WARN_LOG_TAG:  LogTag{name: WARN_LOG_TAG, prefix: "[" + WARN_LOG_TAG + "]"},
}

func getErrorLogTag() LogTag {
	return logTagMap[ERROR_LOG_TAG]
}

func getFatalLogTag() LogTag {
	return logTagMap[FATAL_LOG_TAG]
}

func getInfoLogTag() LogTag {
	return logTagMap[INFO_LOG_TAG]
}

func getPanicLogTag() LogTag {
	return logTagMap[PANIC_LOG_TAG]
}

func getWarnLogTag() LogTag {
	return logTagMap[WARN_LOG_TAG]
}

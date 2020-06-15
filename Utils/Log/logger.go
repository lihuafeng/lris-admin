package Log

import "lris-admin/Config"

func Info(format string, v ...interface{}) {
	logger := logger(Config.LOG_DIR, "")
	logger.setPrefix("INFO")
	logger.logWrite(format+"\n", v...)
}

func Debug(format string, v ...interface{}) {
	logger := logger(Config.LOG_DIR, "")
	logger.setPrefix("DEBUG")
	logger.logWrite(format+"\n", v...)
}
func Warning(format string, v ...interface{}) {
	logger := logger(Config.LOG_DIR, "")
	logger.setPrefix("WARNING")
	logger.logWrite(format+"\n", v...)
}
func Error(format string, v ...interface{}) {
	logger := logger(Config.LOG_DIR, "")
	logger.setPrefix("ERROR")
	logger.logWrite(format+"\n", v...)
}
func Access(format string, v ...interface{}) {
	logger := logger(Config.LOG_DIR, "_access")
	logger.setPrefix("")
	logger.logWrite(format+"\n", v...)
}

package Log

import (
	"log"
	. "lris-admin/Utils/Helper"
	"os"
	"time"
)

type logObj struct {
	logger *log.Logger
}

func (lo logObj) setPrefix(prefix string) {
	if prefix != "" {
		lo.logger.SetPrefix("[" + prefix + "]")
	}
}

func (lo logObj) logWrite(format string, v ...interface{}) {
	lo.logger.Printf(format+"\n", v...)
}

//log对象
func logger(logDir, suffix string) *logObj {
	fileDir, filePath := getFileSuf(logDir, suffix)
	MkdirPath(fileDir)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logObj := &logObj{
		logger: log.New(file, "", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
	return logObj
}

//获取文件存储目录和地址
func getFileSuf(dir, suffix string) (string, string) {
	fileDir := dir + "/" + time.Now().Format("2006-01")
	filePath := fileDir + "/" + time.Now().Format("2006-01-02") + suffix + ".log"
	return fileDir, filePath
}

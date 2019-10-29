package logger

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/CRL-Studio/AuthServer/src/utils/config"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
)

func init() {
	logger()
}

// Log return log
func Log() *zap.SugaredLogger {
	if log == nil {
		logger()
	}
	return log
}

func logger() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	var logPath string
	logPath = fmt.Sprintf("%s/../../../%s%s", dir, config.Get("log.path"), config.Get("log.file"))

	logBuilder := zap.NewDevelopmentConfig()
	logBuilder.OutputPaths = []string{
		logPath,
	}

	logInstance, err := logBuilder.Build()
	if err != nil {
		panic("can't initialize zap logger:" + err.Error())
	}
	log = logInstance.Sugar()
}

// Close Log
func Close() {
	if log != nil {
		log.Sync()
	}
}

package internal

import (
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

var l, _ = os.OpenFile(path.Join(LogPatg, "spider.log"), os.O_WRONLY|os.O_CREATE, 0755)

var SLogger = logger()

func logger() *logrus.Logger {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.MultiWriter(l))
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

package inlogrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Out = os.Stdout

	//logger.WithFields(logrus.Fields{}).Info()
	return logger
}

func WithFileSplite(logger *logrus.Logger) *logrus.Logger {
	if logger == nil {
		return nil
	}
	logger.AddHook(newLfsHook(12))
	return logger
}

func WithElastic(logger *logrus.Logger) *logrus.Logger {
	if logger == nil {
		return nil
	}
	logger.AddHook(Elastic())
	return logger
}

////DefaultFieldHook 默认hook
//type DefaultFieldHook struct {
//}
//
//func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
//	entry.Data["appName"] = "MyAppName"
//	return nil
//}
//
//func (hook *DefaultFieldHook) Levels() []logrus.Level {
//	return logrus.AllLevels
//}

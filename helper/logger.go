package helper

import (
	"achilles/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger initializes the global logger based on the provided LogConfig.
func initializeLogger(logConfig model.LogConfig) {
	globalLogger = createLogger(logConfig)
	globalLogger.Info("Config and logger are loaded.")
}

// GetGlobalLogger returns the globally accessible logger.
func GetGlobalLogger() *logrus.Logger {
	return globalLogger
}

func createLogger(logConfig model.LogConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	if !logConfig.RedirectLogFileToTerminal {
		logger.Formatter = &logrus.JSONFormatter{}
		logger.Out = getLumberjackWriter(&logConfig)
	} else {
		logger.Formatter = &DevErrLogFormatter{}
		logger.Out = os.Stderr
	}

	return logger
}

func getLumberjackWriter(logConfig *model.LogConfig) io.Writer {
	logFilePath := fmt.Sprintf("%s%s.log", logConfig.LogFilePath, time.Now().Format("20060102"))

	return &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    logConfig.LogMaxSizeMB,
		MaxBackups: logConfig.LogMaxBackups,
		MaxAge:     logConfig.LogMaxAgeDays,
	}
}

// DevErrLogFormatter formats logs into pretty-printed JSON.
type DevErrLogFormatter struct{}

// Format renders a single log entry.
func (f *DevErrLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := formatFields(entry)
	data["time"] = entry.Time.Format("2006/01/02 - 15:04:05")
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message

	serialized, err := json.MarshalIndent(data, "  ", "  ")
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return serialized, nil
}

func formatFields(entry *logrus.Entry) logrus.Fields {
	data := make(logrus.Fields, len(entry.Data)+3)
	for key, value := range entry.Data {
		switch v := value.(type) {
		case error:
			data[key] = v.Error()
		default:
			data[key] = v
		}
	}
	return data
}

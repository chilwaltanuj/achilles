package client

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

// BuildAndGetLogWriter creates and returns a logrus.Logger with the specified LogConfig.
func BuildAndGetLogWriter(logConfig model.LogConfig) *logrus.Logger {
	logger := createLogger(logConfig)
	logger.Info("Config and logger are loaded.")
	return logger
}

// createLogger creates a logrus.Logger based on the provided LogConfig.
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

// getLumberjackWriter creates a Lumberjack writer based on LogConfig.
func getLumberjackWriter(logConfig *model.LogConfig) io.Writer {
	logFilePath := fmt.Sprintf("logs/%s-%s.log", logConfig.LogFilePath, time.Now().Format("20060102"))

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
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Serialize errors as strings to include them in JSON
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data["time"] = entry.Time.Format("2006/01/02 - 15:04:05")
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message
	prefix := []byte("")
	serialized, err := json.MarshalIndent(data, "  ", "  ")
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return append(append(prefix, serialized...), '\n'), nil
}

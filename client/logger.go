package client

import (
	"achilles/constant"
	"achilles/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogLevel(logConfigLevel string) logrus.Level {
	var level logrus.Level
	switch logConfigLevel {
	case "panic":
		level = logrus.PanicLevel
	case "fatal":
		level = logrus.FatalLevel
	case "error":
		level = logrus.ErrorLevel
	case "warning":
		level = logrus.WarnLevel
	case "info":
		level = logrus.InfoLevel
	case "debug":
		level = logrus.DebugLevel
	case "trace":
		level = logrus.TraceLevel
	default:
		level = logrus.InfoLevel // Default level is Info if config is not set or invalid
	}
	return level
}

// createLogger creates and configures a new logrus logger.
func BuildAndGetLogger(logConfig model.LogConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(GetLogLevel(logConfig.LogLevel))
	logger.Formatter = &CustomFormatter{logConfig.EnableIndentation}

	if logConfig.LogInTerminalOverFile {
		logger.Out = os.Stderr
	} else {
		logger.Out = getLumberjackWriter(&logConfig)
	}

	return logger
}

// getLumberjackWriter creates a lumberjack log file writer based on LogConfig.
func getLumberjackWriter(logConfig *model.LogConfig) io.Writer {
	logFilePath := fmt.Sprintf("%s%s.log", logConfig.LogFilePath, time.Now().Format("20060102"))
	return &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    logConfig.LogMaxSizeMB,
		MaxBackups: logConfig.LogMaxBackups,
		MaxAge:     logConfig.LogMaxAgeDays,
	}
}

// CustomFormatter is a custom Logrus formatter to handle detailed structured logging.
type CustomFormatter struct {
	EnableIndentation bool
}

// Format formats the log entry for structured data.
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := formatFields(entry)
	logEntry := map[string]any{
		constant.Level:   entry.Level.String(),
		constant.Time:    entry.Time.Format(constant.TimeFormat),
		constant.Message: entry.Message,
		constant.Value:   data,
	}
	var serialized []byte
	var err error
	if f.EnableIndentation {
		serialized, err = json.MarshalIndent(logEntry, "", "  ")
	} else {
		serialized, err = json.Marshal(logEntry)
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	serialized = append(serialized, '\n')
	return serialized, nil
}

// formatFields processes logrus fields, including any errors.
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

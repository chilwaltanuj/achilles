package helper

import (
	"achilles/client"
	"achilles/constant"
	"reflect"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// LogDetails logs a message with structured data.
func LogDetails(logLevelInfo string, message string, dataToLog any) {
	entry := globalLogger.WithFields(logrus.Fields{})
	entry = LogStructFields(entry, dataToLog)
	logrusLogLevel := client.GetLogLevel(logLevelInfo)
	entry.Log(logrusLogLevel, message)
}

func LogMessageWithStackTrace(errorMessage string) {
	LogDetails(constant.LogLevelError, errorMessage, string(debug.Stack()))
}

// LogStructFields logs the fields of a struct in a Logrus entry.
func LogStructFields(entry *logrus.Entry, data any) *logrus.Entry {
	if err, ok := data.(error); ok {
		return entry.WithField("error", err.Error())
	} else if _, ok := data.(string); ok {
		return entry.WithField("text", data)
	}
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Struct {
		fields := logrus.Fields{}
		for index := 0; index < value.NumField(); index++ {
			field := value.Type().Field(index)
			value := value.Field(index).Interface()
			fields[field.Name] = value
		}
		return entry.WithFields(fields)
	}
	return entry
}

package helper

import (
	"achilles/client"
	"reflect"

	"github.com/sirupsen/logrus"
)

// LogDetails logs a message with structured data.
func LogDetails(logLevelInfo string, message string, dataToLog ...any) {
	entry := globalLogger.WithFields(logrus.Fields{})
	entry = LogStructFields(entry, dataToLog...)
	logrusLogLevel := client.GetLogLevel(logLevelInfo)

	entry.Log(logrusLogLevel, message)
}

// LogStructFields logs the fields of a struct in a Logrus entry.
func LogStructFields(entry *logrus.Entry, dataToLog ...any) *logrus.Entry {
	for _, data := range dataToLog {
		if err, ok := data.(error); ok {
			entry = entry.WithField("error", err.Error())
		} else if str, ok := data.(string); ok {
			entry = entry.WithField("text", str)
		} else if mapData, ok := data.(map[string]interface{}); ok {
			entry = entry.WithFields(logrus.Fields(mapData))
		} else {
			value := reflect.ValueOf(data)
			if value.Kind() == reflect.Ptr {
				value = value.Elem() // Dereference the pointer to get the actual value
			}

			if value.Kind() == reflect.Struct {
				fields := logrus.Fields{}
				for index := 0; index < value.NumField(); index++ {
					field := value.Type().Field(index)
					fieldValue := value.Field(index).Interface()
					fields[field.Name] = fieldValue
				}
				entry = entry.WithFields(fields)
			}
		}
	}

	return entry
}

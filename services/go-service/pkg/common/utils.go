package common

import (
	nucliozap "github.com/nuclio/zap"
	"os"
	"reflect"
	"strconv"
)

func ResolveLogLevel(level string) nucliozap.Level {
	switch level {
	case "debug":
		return nucliozap.DebugLevel
	case "info":
		return nucliozap.InfoLevel
	case "warn":
		return nucliozap.WarnLevel
	case "error":
		return nucliozap.ErrorLevel
	default:
		return nucliozap.InfoLevel
	}
}

func GetEnvOrDefault[T any](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	} else if value == "nil" || value == "none" {
		var zeroValue T
		return zeroValue
	}

	switch reflect.TypeOf(defaultValue).Kind() {
	case reflect.String:
		return any(value).(T)
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return any(intValue).(T)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return any(boolValue).(T)
	default:
		return defaultValue
	}
}

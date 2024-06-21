package utils

import "log/slog"

func LogError(method string, message string, err error, other ...string) {
	if err != nil {
		slog.Error("method", method, "msg", message, err, other)
	} else {
		slog.Error("method", method, "msg", message, other)
	}
}

func LogInfo(method string, message string, key any) {
	slog.Info("method", method, "msg", message, key)
}

func LogDebug(method string, message string, key string) {
	slog.Debug("method", method, "msg", message, key)
}

func LogWarn(method string, message string, extra any) {
	slog.Warn("method", method, "msg", message, extra)
}

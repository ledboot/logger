package logger

import (
	"testing"
)

func show() {
	Debug("debug message")
	Debugf("debugf message: %s", "test...")
	Debugw("test...", "key", "value")

	Info("info message")
	Infof("infof message: %s", "test...")
	Infow("test...", "key", "value")

	Warn("warn message")
	Warnf("warnf message: %s", "test...")
	Warnw("test...", "key", "value")

	Error("Error message")
	Errorf("Errorf message: %s", "test...")
	Errorw("test...", "key", "value")
}

func TestDefultLogger(t *testing.T) {
	show()
	Info("------------------")
	var l Level
	_ = l.Set("error")
	SetLevel(l)
	show()
}

func TestLevel(t *testing.T) {
	show()
}

func TestProdLogger(t *testing.T) {
	config := NewProductionConfig(FieldPair{"service", "client_string"}, FieldPair{"version", "v0.1.0-5309251"})
	SetConfig(config)
	var l Level
	_ = l.Set("debug")
	SetLevel(l)

	With("trace_id", "274ac2bbf9d5")
	With("span_id", "383d60f1")

	show()
}

func TestDevFileLogger(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = DebugLevel
	config.OutputPaths = []string{"stdout", "dev.log", "test.log"}
	config.DisableStacktrace = true
	config.DisableCaller = false
	config.ShortTime = true
	config.EnableColor = false
	SetConfig(config)

	show()
}

func TestColorLogger(t *testing.T) {
	config := NewDevelopmentConfig()
	config.Level = DebugLevel
	config.DisableStacktrace = true
	config.DisableCaller = true
	SetConfig(config)

	show()
}

func TestLoggerInterface(t *testing.T) {
	l := GetLogger()
	l.Info("get logger...")
}

package logger

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Warn(args ...interface{}) {
	l.Warn(args...)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func DPanic(args ...interface{}) {
	l.DPanic(args...)
}

func Panic(args ...interface{}) {
	l.Panic(args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	l.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	l.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	l.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	l.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	l.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	l.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	l.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	l.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	l.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	l.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	l.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	l.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	l.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	l.Fatalw(msg, keysAndValues...)
}

func Sync() error {
	return l.Sync()
}

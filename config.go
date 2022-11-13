package logger

import "go.uber.org/zap"

type FieldPair []string

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level Level `json:"level" yaml:"level"`

	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `json:"development" yaml:"development"`

	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`

	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`

	// Encoding sets the logger's encoding. Valid values are "json" and "console"
	Encoding string `json:"encoding" yaml:"encoding"`

	// OutputPaths is a list of URLs or file paths to write logging output to.
	// See Open for details.
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`

	// InitialFields is a collection of fields to add to the root logger.
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`

	EnableColor bool
	ShortTime   bool

	CallerSkip int
	zapConfig  *zap.Config
}

func NewProductionConfig(fields ...FieldPair) *Config {
	return &Config{
		Level:             InfoLevel,
		Development:       false,
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		CallerSkip:        2,
		DisableStacktrace: false,
		InitialFields:     genInitialFields(fields),
	}
}

func NewDevelopmentConfig(fields ...FieldPair) *Config {
	return &Config{
		Level:             DebugLevel,
		Development:       true,
		ShortTime:         true,
		EnableColor:       true,
		Encoding:          "console",
		OutputPaths:       []string{"stdout"},
		CallerSkip:        2,
		DisableStacktrace: true,
		InitialFields:     genInitialFields(fields),
	}
}

func NewDefaultConfig(fields ...FieldPair) *Config {
	return NewProductionConfig(fields...)
}

func genInitialFields(args []FieldPair) map[string]interface{} {
	fields := make(map[string]interface{})
	for _, f := range args {
		fields[f[0]] = f[1]
	}
	return fields
}

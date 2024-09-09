package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey int

const (
	loggerContextKey contextKey = iota
)

var (
	// global logger instance.
	global       *zap.SugaredLogger
	currentLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
)

// commonly used key-values for logger decoration
var (
	LongTermArgs   = []interface{}{"@lt", ""}
	generationArgs = []interface{}{"@gen", "1"}
)

func init() { //nolint
	SetLogger(New())
}

// WithUniqueFields keep fields unique.
func WithUniqueFields() zap.Option {
	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return &uniqueFieldsCore{core, nil}
	})
}

// New creates new *zap.SugaredLogger with standard EncoderConfig
// if lvl == nil, global AtomicLevel will be used
func New(options ...zap.Option) *zap.SugaredLogger {
	return NewWithSink(os.Stdout, options...)
}

// NewWithSink ...
func NewWithSink(sink io.Writer, options ...zap.Option) *zap.SugaredLogger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(sink),
			currentLevel,
		),
		options...,
	).Sugar().With(generationArgs...)
}

// Level returns current global logger level
func Level() zapcore.Level {
	return currentLevel.Level()
}

// SetLevel sets level for global logger
func SetLevel(l zapcore.Level) {
	currentLevel.SetLevel(l)
	SetLogger(New())
}

// Logger returns current global logger.
func Logger() *zap.SugaredLogger {
	return global
}

// SetLogger sets global used logger. This function is not thread-safe.
func SetLogger(l *zap.SugaredLogger) {
	global = l
}

// FromContext returns logger from context if set
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := global

	if logger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		l = logger
	}

	return l
}

// Debug ...
func Debug(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Debug(args...)
}

// Debugf ...
func Debugf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Debugf(format, args...)
}

// DebugKV ...
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Debugw(message, kvs...)
}

// Info ...
func Info(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Info(args...)
}

// Infof ...
func Infof(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Infof(format, args...)
}

// InfoKV ...
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Infow(message, kvs...)
}

// Warn ...
func Warn(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Warn(args...)
}

// Warnf ...
func Warnf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Warnf(format, args...)
}

// WarnKV ...
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Warnw(message, kvs...)
}

// Error ...
func Error(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Error(args...)
}

// Errorf ...
func Errorf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Errorf(format, args...)
}

// ErrorKV ...
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Errorw(message, kvs...)
}

// Fatal ...
func Fatal(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Fatal(args...)
}

// Fatalf ...
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Fatalf(format, args...)
}

// FatalKV ...
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Fatalw(message, kvs...)
}

type uniqueFieldsCore struct {
	zapcore.Core
	fieldKeys map[string]bool
}

func (c *uniqueFieldsCore) With(fields []zapcore.Field) zapcore.Core {
	filtered := make([]zapcore.Field, 0, len(fields))

	fieldKeys := make(map[string]bool, len(c.fieldKeys))
	for k, v := range c.fieldKeys {
		fieldKeys[k] = v
	}

	for i := range fields {
		if fieldKeys[fields[i].Key] {
			continue
		}
		filtered = append(filtered, fields[i])
		fieldKeys[fields[i].Key] = true
	}

	return &uniqueFieldsCore{c.Core.With(filtered), fieldKeys}
}

func LowercaseLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	level := ""
	switch l {
	case zapcore.ErrorLevel:
		level = "error"
	case zapcore.WarnLevel:
		level = "warn"
	default:
		level = l.String()
	}
	enc.AppendString(level)
}

// logrus.go
// 实现logger来包装logrus，并实现Logger接口

package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var _ Logger = (*logger)(nil)

type logger struct {
	e *logrus.Entry
}

func NewLogger(w io.Writer, options ...Options) Logger {
	l := logrus.New()
	l.Out = w
	l.Level = logrus.InfoLevel
	l.Formatter = &logrus.TextFormatter{}

	for _, option := range options {
		option(l)
	}

	entry := logrus.NewEntry(l)
	return &logger{
		e: entry,
	}
}

// logrus Fields
type Fields map[string]interface{}

func WithFields(l Logger, fields Fields) Logger {
	f := logrus.Fields(fields)
	l2, ok := l.(*logger)
	if !ok {
		return l
	}

	return l2.e.WithFields(f)
}

// logrus options
type Options func(*logrus.Logger)

// log level option
// TODO: logrus提供了Fatal等level
// 但我的接口是不提供的
// 虽说自己用不会雍错，但这里需不需要做点操作来避免？
func LogLevel(lvl logrus.Level) Options {
	return func(l *logrus.Logger) {
		l.Level = lvl
	}
}

// logrus formatter option
func LogFormat(formatter logrus.Formatter) Options {
	return func(l *logrus.Logger) {
		l.Formatter = formatter
	}
}

func (l *logger) Debug(arg ...interface{}) {
	l.e.Debug(arg...)
}

func (l *logger) Info(arg ...interface{}) {
	l.e.Info(arg...)
}

func (l *logger) Warning(arg ...interface{}) {
	l.e.Warning(arg...)
}

func (l *logger) Error(arg ...interface{}) {
	l.e.Error(arg...)
}

func (l *logger) Debugf(format string, arg ...interface{}) {
	l.e.Debugf(format, arg...)
}

func (l *logger) Infof(format string, arg ...interface{}) {
	l.e.Infof(format, arg...)
}

func (l *logger) Warningf(format string, arg ...interface{}) {
	l.e.Warningf(format, arg...)
}

func (l *logger) Errorf(format string, arg ...interface{}) {
	l.e.Errorf(format, arg...)
}

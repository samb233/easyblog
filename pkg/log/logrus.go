// logrus.go
// 实现logger来包装logrus，并实现Logger接口

package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var _ Logger = (*logger)(nil)

type logger struct {
	e *logrus.Entry
}

func NewLogger(w io.Writer, options ...Options) Logger {
	l := logrus.New()
	l.Out = os.Stdout
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

// 实现Fields
type Fields map[string]interface{}

func WithFields(l Logger, fields Fields) Logger {
	f := logrus.Fields(fields)
	l2, ok := l.(*logger)
	if !ok {
		return l
	}

	return l2.e.WithFields(f)
}

// logrus配置设置
type Options func(*logrus.Logger)

// 日志等级设置
func LogLevel(lvl logrus.Level) Options {
	return func(l *logrus.Logger) {
		l.Level = lvl
	}
}

// 输出格式配置
func LogFormat(formatter logrus.Formatter) Options {
	return func(l *logrus.Logger) {
		l.Formatter = formatter
	}
}

// 实现Logger接口
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

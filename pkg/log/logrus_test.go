package log

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	tests := map[string]struct {
		level     logrus.Level
		formatter logrus.Formatter
		message   string
		want      string
	}{
		"json format": {
			level:     logrus.DebugLevel,
			formatter: &logrus.JSONFormatter{},
			message:   "this is just a test",
			want:      `{"level":"debug","msg":"this is just a test"`,
		},
		"unmatch level": {
			level:     logrus.InfoLevel,
			formatter: &logrus.JSONFormatter{},
			message:   "this is just a test",
			want:      "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			output := new(bytes.Buffer)
			logger := NewLogger(output, LogLevel(test.level), LogFormat(test.formatter))
			logger.Debug(test.message)
			// fmt.Println(output.String())
			// fmt.Println(test.want)
			assert.True(t, strings.HasPrefix(output.String(), test.want))
		})
	}
}

func TestWithFields(t *testing.T) {
	output := new(bytes.Buffer)
	logger := NewLogger(output, LogLevel(logrus.InfoLevel), LogFormat(&logrus.JSONFormatter{}))
	logger = WithFields(logger, Fields{
		"name": "log",
	})
	message := "this is just a test"

	t.Run("WithFields once", func(t *testing.T) {
		t.Helper()
		output.Reset()
		want := `{"level":"info","msg":"this is just a test","name":"log"`

		logger.Info(message)
		assert.True(t, strings.HasPrefix(output.String(), want))
	})

	t.Run("WithFields twice", func(t *testing.T) {
		t.Helper()
		output.Reset()
		loggerTwice := WithFields(logger, Fields{
			"test": "this is just a test",
		})
		want := `{"level":"info","msg":"this is just a test","name":"log","test":"this is just a test"`

		loggerTwice.Info(message)
		// fmt.Println(output.String())
		assert.True(t, strings.HasPrefix(output.String(), want))
	})

}

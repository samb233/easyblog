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

package golangruntime

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

// UTCFormatter ...
type UTCFormatter struct {
	easy.Formatter
}

// Format ...
func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

// GenLoggerByRequestID ...
func GenLoggerByRequestID(rid string) *logrus.Entry {
	l := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &UTCFormatter{
			easy.Formatter{
				TimestampFormat: "2006-01-02T15:04:05.999Z",
				LogFormat:       "%time%: %requestId% [%lvl%]  %msg%\n",
			},
		},
	}
	le2 := l.WithField("requestId", rid)
	return le2
}

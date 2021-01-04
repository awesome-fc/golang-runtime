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

var log = logrus.New()

var logMap map[string]*logrus.Entry

func initLogger() {
	log = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &UTCFormatter{
			easy.Formatter{
				TimestampFormat: "2006-01-02T15:04:05.999Z",
				LogFormat:       "%time%: %requestId% [%lvl%]  %msg%\n",
			},
		},
	}
	logMap = make(map[string]*logrus.Entry)
}

// GetLogger ...
func GetLogger() *logrus.Logger {
	return log
}

// GetLoggerByRequestID ...
func GetLoggerByRequestID(rid string) *logrus.Entry {
	if le, ok := logMap[rid]; ok {
		return le
	}
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
	logMap[rid] = le2
	return le2
}

// SetLoggerLevel ...
func SetLoggerLevel(level logrus.Level) *logrus.Logger {
	log = &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Formatter: &UTCFormatter{
			easy.Formatter{
				TimestampFormat: "2006-01-02T15:04:05.999Z",
				LogFormat:       "%time%: %requestId% [%lvl%]  %msg%\n",
			},
		},
	}
	return log
}

package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Log refers to the logger instance used by the permissions service.
var Log = logrus.WithFields(logrus.Fields{
	"service": "permissions",
	"art-id":  "permissions",
	"group":   "org.cyverse",
})

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// LogClose closes c and logs any returned error using the package logger. The
// name is included in the log message to identify what was being closed.
func LogClose(c io.Closer, name string) {
	if err := c.Close(); err != nil {
		Log.Errorf("closing %s: %v", name, err)
	}
}

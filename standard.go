package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	datetimeFormat = "2006-01-02 15:04:05 "
)

var (
	// DefaultLogger is a Logger that writes datetime prefixed log messages to stdout.
	DefaultLogger = &StandardLogger{
		Writer: os.Stdout,
	}

	// ErrorLogger is a Logger that writes datetime prefixed log messages to stderr.
	ErrorLogger = &StandardLogger{
		Writer: os.Stdout,
	}

	// DiscardLogger is a Logger that discards all log messages writtern to it.
	DiscardLogger = &StandardLogger{
		Writer: io.Discard,
	}
)

// StandardLogger is a Logger that writes datetime prefixed log messages to an io.Writer.
type StandardLogger struct {
	Writer io.Writer
}

// Log writes a log message to the underlying io.Writer. Arguments are handled in the manner of fmt.Print.
func (l *StandardLogger) Log(v ...interface{}) {
	b := []byte(fmt.Sprint(v...))
	if _, err := l.Write(b); err != nil {
		panic(err)
	}
}

// Logf writes a log message to the underlying io.Writer. Arguments are handled in the manner of fmt.Printf.
func (l *StandardLogger) Logf(format string, v ...interface{}) {
	b := []byte(fmt.Sprintf(format, v...))
	if _, err := l.Write(b); err != nil {
		panic(err)
	}
}

// Write writes a log message to the underlying io.Writer and returns the number of bytes written, or error on failure.
func (l *StandardLogger) Write(b []byte) (n int, err error) {
	buf := []byte{}
	buf = append(buf, time.Now().Format(datetimeFormat)...)
	buf = append(buf, b...)
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	return l.Writer.Write(buf)
}

package log

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Package defaults
var (
	DefaultTimeFormat = "2006-01-02 15:04:05 "
	DefaultPrefixFunc = DateTimePrefix(DefaultTimeFormat)
	DefaultWriter     = os.Stderr
	DefaultLogger     = &Logger{
		Writer: DefaultWriter,
		Prefix: DefaultPrefixFunc,
	}
)

// Logger is responsible for writing log messages.
type Logger struct {
	Writer io.Writer
	Prefix PrefixFunc
}

// PrefixFunc is a function responsible for returning a log prefix.
type PrefixFunc func() string

// DateTimePrefix produces a PrefixFunc which returns the current date & time as a prefix.
func DateTimePrefix(format string) PrefixFunc {
	return func() string {
		return time.Now().Format(format)
	}
}

// Log writes a log message via the underlying Writer. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Log(v ...interface{}) {
	b := []byte(fmt.Sprint(v...))
	if _, err := l.Write(b); err != nil {
		panic(err)
	}
}

// Logf writes a log message via the underlying Writer. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Logf(format string, v ...interface{}) {
	b := []byte(fmt.Sprintf(format, v...))
	if _, err := l.Write(b); err != nil {
		panic(err)
	}
}

// Write writes a log message via the underlying Writer and returns the number of bytes written, or error on failure.
func (l *Logger) Write(b []byte) (n int, err error) {
	buf := []byte(l.Prefix())
	buf = append(buf, b...)
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	return l.Writer.Write(buf)
}

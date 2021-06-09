package log

// Logger is an interface for writing log messages.
type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
	Write(b []byte) (n int, err error)
}

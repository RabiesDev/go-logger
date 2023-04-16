package log

import "fmt"

func (log *Logger) Warn(args ...interface{}) {
	printer := func() error {
		log.ApplyPrefix(log.Prefixes.Warn)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	}
	_ = log.Print(printer)
}

func (log *Logger) Warnf(format string, args ...interface{}) {
	printer := func() error {
		log.ApplyPrefix(log.Prefixes.Warn)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	}
	_ = log.Print(printer)
}

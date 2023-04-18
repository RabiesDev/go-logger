package log

import "fmt"

func (log *Logger) Error(args ...interface{}) {
	if log.Level > 3 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Error)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

func (log *Logger) Errorf(format string, args ...interface{}) {
	if log.Level > 3 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Error)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

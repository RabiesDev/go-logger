package log

import "fmt"

func (log *Logger) Warn(args ...interface{}) {
	if log.Level > 3 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Warn)
		_, err := fmt.Fprint(log.buffer, args...)
		return err
	})
}

func (log *Logger) Warnln(args ...interface{}) {
	if log.Level > 3 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Warn)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

func (log *Logger) Warnf(format string, args ...interface{}) {
	if log.Level > 3 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Warn)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

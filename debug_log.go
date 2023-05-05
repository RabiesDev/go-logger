package log

import "fmt"

func (log *Logger) Debug(args ...interface{}) {
	if log.Level > 1 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Debug)
		_, err := fmt.Fprint(log.buffer, args...)
		return err
	})
}

func (log *Logger) Debugln(args ...interface{}) {
	if log.Level > 1 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Debug)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

func (log *Logger) Debugf(format string, args ...interface{}) {
	if log.Level > 1 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Debug)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

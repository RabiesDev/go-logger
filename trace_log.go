package log

import "fmt"

func (log *Logger) Trace(args ...interface{}) {
	if log.Level > 0 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Trace)
		_, err := fmt.Fprint(log.buffer, args...)
		return err
	})
}

func (log *Logger) Traceln(args ...interface{}) {
	if log.Level > 0 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Trace)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

func (log *Logger) Tracef(format string, args ...interface{}) {
	if log.Level > 0 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Trace)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

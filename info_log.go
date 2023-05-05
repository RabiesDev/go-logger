package log

import "fmt"

func (log *Logger) Info(args ...interface{}) {
	if log.Level > 2 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Info)
		_, err := fmt.Fprint(log.buffer, args...)
		return err
	})
}

func (log *Logger) Infoln(args ...interface{}) {
	if log.Level > 2 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Info)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

func (log *Logger) Infof(format string, args ...interface{}) {
	if log.Level > 2 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Info)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

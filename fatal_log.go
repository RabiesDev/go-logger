package log

import "fmt"

func (log *Logger) Fatal(args ...interface{}) {
	if log.Level > 5 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Fatal)
		_, err := fmt.Fprint(log.buffer, args...)
		return err
	})
}

func (log *Logger) Fatalln(args ...interface{}) {
	if log.Level > 5 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Fatal)
		_, err := fmt.Fprint(log.buffer, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

func (log *Logger) Fatalf(format string, args ...interface{}) {
	if log.Level > 5 {
		return
	}
	_ = log.Print(func() error {
		log.ApplyPrefix(log.Prefixes.Fatal)
		_, err := fmt.Fprintf(log.buffer, format, args...)
		log.buffer.Write([]byte("\n"))
		return err
	})
}

package log

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

var (
	Reset = []byte("\033[0m")

	Bold      = []byte("\\033[1m")
	Underline = []byte("\033[4m")

	Black  = []byte("\\033[30m")
	Red    = []byte("\033[0;31m")
	Green  = []byte("\033[0;32m")
	Yellow = []byte("\033[33m")
	Orange = []byte("\033[0;33m")
	Blue   = []byte("\033[0;34m")
	Purple = []byte("\033[0;35m")
	Cyan   = []byte("\033[0;36m")
	Gray   = []byte("\033[0;37m")
)

type Printer func() error

type Prefix struct {
	Plain []byte
	Color []byte
}

type Prefixes struct {
	Info  Prefix
	Warn  Prefix
	Error Prefix
	Fatal Prefix
	Debug Prefix
}

type Logger struct {
	Prefixes Prefixes
	Out      io.Writer

	buffer *bytes.Buffer
	mutex  *sync.RWMutex
	name   *string
	color  bool
}

func DefaultPrefixes() Prefixes {
	return Prefixes{
		Info: Prefix{
			Plain: []byte("INFO"),
			Color: Cyan,
		},
		Warn: Prefix{
			Plain: []byte("WARN"),
			Color: Yellow,
		},
		Error: Prefix{
			Plain: []byte("ERROR"),
			Color: Red,
		},
		Fatal: Prefix{
			Plain: []byte("FATAL"),
			Color: Orange,
		},
		Debug: Prefix{
			Plain: []byte("DEBUG"),
			Color: Purple,
		},
	}
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{
		Prefixes: DefaultPrefixes(),
		Out:      out,

		buffer: new(bytes.Buffer),
		mutex:  new(sync.RWMutex),
	}
}

func (log *Logger) WithColor() *Logger {
	log.color = true
	return log
}

func (log *Logger) SetName(name string) {
	log.name = &name
}

func (log *Logger) ApplyColor(data []byte, color []byte) []byte {
	if !log.color {
		return data
	}
	var result []byte
	return append(append(append(result, color...), data...), Reset...)
}

func (log *Logger) ApplyPrefix(prefix Prefix) {
	if log.name != nil && len(*log.name) > 0 {
		log.buffer.Write([]byte(fmt.Sprintf("[%s] [%s] ", *log.name, log.ApplyColor(prefix.Plain, prefix.Color))))
	} else {
		log.buffer.Write([]byte(fmt.Sprintf("[%s] ", log.ApplyColor(prefix.Plain, prefix.Color))))
	}
}

func (log *Logger) Println(args ...interface{}) error {
	printer := func() error {
		_, err := fmt.Fprint(log.buffer, args...)
		return err
	}
	return log.Print(printer)
}

func (log *Logger) Printf(format string, args ...interface{}) error {
	printer := func() error {
		_, err := fmt.Fprintf(log.buffer, format, args...)
		return err
	}
	return log.Print(printer)
}

func (log *Logger) Print(writer Printer) error {
	log.mutex.Lock()
	defer log.mutex.Unlock()

	log.buffer.Reset()
	log.buffer.Write(Reset)
	if err := writer(); err != nil {
		return err
	}
	_, err := log.Out.Write(log.buffer.Bytes())
	return err
}

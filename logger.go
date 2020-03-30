package playlistmaker

import (
	"fmt"
	"log"
)

const (
	Debug = iota
	Info  = iota
	Warn  = iota
)

//MyLogger logger of package
type MyLogger struct {
	log.Logger
	level int
}

//SetLevel  message
func (l *MyLogger) SetLevel(level int) {
	l.level = level
}

//Debug debug message
func (l *MyLogger) Debug(v interface{}) {
	if l.level <= Debug {
		fmt.Printf("DEBUG : %#v\n", v)
	}
}

//Info debug message
func (l *MyLogger) Info(v interface{}) {
	if l.level <= Info {
		fmt.Printf("INFO : %#v\n", v)
	}
}

//Warn debug message
func (l *MyLogger) Warn(v interface{}) {
	if l.level <= Warn {
		fmt.Printf("WARN : %#v\n", v)
	}
}

var singletonLogger *MyLogger

//LogInstance get logger Instance
func LogInstance() *MyLogger {
	if singletonLogger == nil {
		singletonLogger = new(MyLogger)
		singletonLogger.SetLevel(Warn)
	}
	return singletonLogger
}

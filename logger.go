package playlistmaker

import (
	"fmt"
	"log"
)

//MyLogger logger of package
type MyLogger struct {
	log.Logger
}

//Debug debug message
func (l *MyLogger) Debug(v interface{}) {
	fmt.Printf("DEBUG : %#v\n", v)
}

var singletonLogger *MyLogger

//LogInstance get logger Instance
func LogInstance() *MyLogger {
	if singletonLogger == nil {
		singletonLogger = new(MyLogger)
	}
	return singletonLogger
}

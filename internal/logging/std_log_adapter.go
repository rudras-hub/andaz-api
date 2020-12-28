package logging

import (
	"fmt"
	"log"
	"os"
)

type logPrefix string

const(
	Debug logPrefix = "DEBUG"
	Info logPrefix = "INFO"
	Warn logPrefix = "WARNING"
	Error logPrefix = "ERROR"

)

type LogExtension struct{
	*log.Logger
}

func NewStdLoggerExtension() *LogExtension{
	stdLogger := log.New(os.Stderr, "", log.LstdFlags)
	return &LogExtension{stdLogger}
}

func (l *LogExtension)Debugf(format string, values ...interface{}){
	l.Println(appendLogPrefix(Debug, format, values))
}

func (l *LogExtension)Infof(format string, values ...interface{}){
	l.Println(appendLogPrefix(Info, format, values))
}
func (l *LogExtension)Warnf(format string, values ...interface{}){
	l.Println(appendLogPrefix(Warn, format, values))
}
func (l *LogExtension)Errorf(format string, values ...interface{}){
	l.Println(appendLogPrefix(Error, format, values))
}

func appendLogPrefix(prefixType logPrefix, originalString string, values ...interface{}) string{
	prefixStr := string(prefixType)
	return fmt.Sprintln(prefixStr, fmt.Sprintf(originalString, values...))
}


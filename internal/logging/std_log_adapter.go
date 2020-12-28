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

func NewStdLoggerExtension(serviceName string) *LogExtension{
	stdLogger := log.New(os.Stdout, serviceName, log.LstdFlags)
	return &LogExtension{stdLogger}
}

func (l *LogExtension)Debug(values ...interface{}){
	l.Println(appendPrefix(Debug, "", values))
}

func (l *LogExtension)Debugf(format string, values ...interface{}){
	l.Println(appendPrefix(Debug, format, values))
}

func (l *LogExtension)Info(values ...interface{}){
	l.Println(appendPrefix(Info, "", values))
}

func (l *LogExtension)Infof(format string, values ...interface{}){
	l.Println(appendPrefix(Info, format, values))
}

func (l *LogExtension)Warn(values ...interface{}){
	l.Println(appendPrefix(Warn, "", values))
}

func (l *LogExtension)Warnf(format string, values ...interface{}){
	l.Println(appendPrefix(Warn, format, values))
}

func (l *LogExtension)Error(values ...interface{}){
	l.Println(appendPrefix(Error, "", values))
}

func (l *LogExtension)Errorf(format string, values ...interface{}){
	l.Println(appendPrefix(Error, format, values))
}

func appendPrefix(prefixType logPrefix, formattedString string, values ...interface{}) string{
	prefixStr := string(prefixType)
	if formattedString == "" {
		return fmt.Sprintln(prefixStr, fmt.Sprint(values...))
	}
	return fmt.Sprintln(prefixStr, fmt.Sprintf(formattedString, values...))
}


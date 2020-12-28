package entity

// Logger interface to provide application logging.
type Logger interface{
	Debug(values ...interface{})
	Debugf(format string, values ...interface{})
	Info(values ...interface{})
	Infof(format string, values ...interface{})
	Warn(values ...interface{})
	Warnf(format string, values ...interface{})
	Error(values ...interface{})
	Errorf(format string, values ...interface{})
}

package logger

type Logger interface {
	WithContext(args ...interface{}) Logger
	Named(name string) Logger

	Info(args ...interface{})
	Debug(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})

	Infow(message string, args ...interface{})
	Debugw(message string, args ...interface{})
	Warningw(message string, args ...interface{})
	Errorw(message string, args ...interface{})
	Panicw(message string, args ...interface{})

	Infof(message string, args ...interface{})
	Debugf(message string, args ...interface{})
	Warningf(message string, args ...interface{})
	Errorf(message string, args ...interface{})
	Panicf(message string, args ...interface{})
}

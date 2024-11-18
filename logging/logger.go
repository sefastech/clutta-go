package logging

import "github.com/sirupsen/logrus"

// Logger defines a generic logging interface
type Logger interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

type logrusLogger struct {
	*logrus.Logger
}

func (l *logrusLogger) Trace(args ...interface{}) { l.Logger.Trace(args...) }
func (l *logrusLogger) Debug(args ...interface{}) { l.Logger.Debug(args...) }
func (l *logrusLogger) Info(args ...interface{})  { l.Logger.Info(args...) }
func (l *logrusLogger) Warn(args ...interface{})  { l.Logger.Warn(args...) }
func (l *logrusLogger) Error(args ...interface{}) { l.Logger.Error(args...) }
func (l *logrusLogger) Fatal(args ...interface{}) { l.Logger.Fatal(args...) }
func (l *logrusLogger) Panic(args ...interface{}) { l.Logger.Panic(args...) }

func (l *logrusLogger) Tracef(format string, args ...interface{}) { l.Logger.Tracef(format, args...) }
func (l *logrusLogger) Debugf(format string, args ...interface{}) { l.Logger.Debugf(format, args...) }
func (l *logrusLogger) Infof(format string, args ...interface{})  { l.Logger.Infof(format, args...) }
func (l *logrusLogger) Warnf(format string, args ...interface{})  { l.Logger.Warnf(format, args...) }
func (l *logrusLogger) Errorf(format string, args ...interface{}) { l.Logger.Errorf(format, args...) }
func (l *logrusLogger) Fatalf(format string, args ...interface{}) { l.Logger.Fatalf(format, args...) }
func (l *logrusLogger) Panicf(format string, args ...interface{}) { l.Logger.Panicf(format, args...) }

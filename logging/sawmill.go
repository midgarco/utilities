package logging

import (
	log "github.com/sirupsen/logrus"
	Lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Fields ...
type Fields map[string]interface{}

// Logger interface
type Logger interface {
	IncludeGlobalFields(f Fields)

	WithField(key string, value interface{}) *log.Entry
	WithFields(fields Fields) *log.Entry
	WithError(err error) *log.Entry

	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Print(args ...interface{})
	Panic(args ...interface{})

	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

// Sawmill ...
type Sawmill struct {
	logger *log.Logger
	fields Fields
}

// NewFileLogger ...
func NewFileLogger(filename string, size, age int) *Sawmill {
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = &Lumberjack.Logger{
		Filename: filename,
		MaxSize:  size,
		MaxAge:   age,
	}
	return &Sawmill{logger, Fields{}}
}

// NewLogger ...
func NewLogger() *Sawmill {
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	return &Sawmill{logger, Fields{}}
}

// SetLevel ...
func (l *Sawmill) SetLevel(level string) {
	ll, err := log.ParseLevel(level)
	if err != nil {
		ll = log.InfoLevel
	}
	l.logger.SetLevel(ll)
}

// IncludeGlobalFields ...
func (l *Sawmill) IncludeGlobalFields(f Fields) {
	for k, v := range f {
		l.fields[k] = v
	}
}

// WithField ...
func (l Sawmill) WithField(key string, value interface{}) *log.Entry {
	f := log.Fields{key: value}
	for k, v := range l.fields {
		f[k] = v
	}
	return l.logger.WithFields(log.Fields(f))
}

// WithFields ...
func (l Sawmill) WithFields(f log.Fields) *log.Entry {
	for k, v := range l.fields {
		f[k] = v
	}
	return l.logger.WithFields(log.Fields(f))
}

// WithError ...
func (l Sawmill) WithError(err error) *log.Entry {
	return l.logger.WithError(err)
}

// Info ...
func (l Sawmill) Info(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Info(args...)
	} else {
		l.logger.Info(args...)
	}
}

// Infof ...
func (l Sawmill) Infof(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Infof(format, args...)
	} else {
		l.logger.Infof(format, args...)
	}
}

// Debug ...
func (l Sawmill) Debug(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Debug(args...)
	} else {
		l.logger.Debug(args...)
	}
}

// Debugf ...
func (l Sawmill) Debugf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Debugf(format, args...)
	} else {
		l.logger.Debugf(format, args...)
	}
}

// Warn ...
func (l Sawmill) Warn(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Warn(args...)
	} else {
		l.logger.Warn(args...)
	}
}

// Warnf ...
func (l Sawmill) Warnf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Warnf(format, args...)
	} else {
		l.logger.Warnf(format, args...)
	}
}

// Warning ...
func (l Sawmill) Warning(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Warning(args...)
	} else {
		l.logger.Warning(args...)
	}
}

// Warningf ...
func (l Sawmill) Warningf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Warningf(format, args...)
	} else {
		l.logger.Warningf(format, args...)
	}
}

// Error ...
func (l Sawmill) Error(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Error(args...)
	} else {
		l.logger.Error(args...)
	}
}

// Errorf ...
func (l Sawmill) Errorf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Errorf(format, args...)
	} else {
		l.logger.Errorf(format, args...)
	}
}

// Fatal ...
func (l Sawmill) Fatal(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Fatal(args...)
	} else {
		l.logger.Fatal(args...)
	}
}

// Fatalf ...
func (l Sawmill) Fatalf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Fatalf(format, args...)
	} else {
		l.logger.Fatalf(format, args...)
	}
}

// Print ...
func (l Sawmill) Print(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Print(args...)
	} else {
		l.logger.Print(args...)
	}
}

// Printf ...
func (l Sawmill) Printf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Printf(format, args...)
	} else {
		l.logger.Printf(format, args...)
	}
}

// Panic ...
func (l Sawmill) Panic(args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Panic(args...)
	} else {
		l.logger.Panic(args...)
	}
}

// Panicf ...
func (l Sawmill) Panicf(format string, args ...interface{}) {
	if len(l.fields) > 0 {
		l.logger.WithFields(log.Fields(l.fields)).Panicf(format, args...)
	} else {
		l.logger.Panicf(format, args...)
	}
}

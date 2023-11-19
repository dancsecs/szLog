/*
   Szerszam Log Utility: szLog.
   Copyright (C) 2023  Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
Package szLog provides for writing to logs with four levels of detail as
follows:

- Debug
- Info
- Warn
- Error

It layers on top of the standard golang log package following its design lead
providing for both a default (standard) logger that can be directly accesses
with package level functions and variables or can create an independent
logging object to be used by applications.  Multiple log.Loggers can be added
as long as they reference different underlying io.Writer objects and each can
have its own flags.
*/
//nolint:goCheckNoGlobals,goCheckNoInits // ok
package szLog

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// Level stores the current level of permitted logging.
type Level uint32

// Defines the various logging levels.
const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

const (
	debugLabel    = "D: "
	infoLabel     = "I: "
	warnLabel     = "W: "
	errorLabel    = "E: "
	continueLabel = "+  "
)

// Logger represents a szLog logging object.
type Logger struct {
	level   Level
	logs    []*log.Logger
	IsDebug bool
	IsInfo  bool
	IsWarn  bool
}

// New creats a new szLog.Logger with the provided logging level.
func New(level Level, logLogger *log.Logger) *Logger {
	logger := new(Logger)
	logger.SetLevel(level)
	_ = logger.AddLogger(logLogger)
	return logger
}

// Output writes the labeled stecified message to all szLog.Loggers added.
func (logger *Logger) output(label, msg string) {
	r := label
	for i, l := range strings.Split(msg, "\n") {
		if i > 0 {
			r += "\n" + continueLabel
		}
		r += l
	}
	for _, l := range logger.logs {
		l.Print(r)
	}
}

// SetLevel sets the logging level for the Logger.
func (logger *Logger) SetLevel(newLevel Level) Level {
	lastLevel := logger.level
	logger.level = newLevel
	logger.IsWarn = logger.level >= WarnLevel
	logger.IsInfo = logger.level >= InfoLevel
	logger.IsDebug = logger.level >= DebugLevel
	return lastLevel
}

// AddWriter wraps the provided io.Writer in a new log.Logger and adds it
// to the logs output by the selected szLog.Logger.
func (logger *Logger) AddWriter(
	newWriter io.Writer, prefix string, flags int,
) error {
	return logger.AddLogger(log.New(newWriter, prefix, flags))
}

// AddLogger adds the provided log.Logger to the logs output by the
// selected szLog.Logger.  Checks are made and an error is returned should
// duplicate szLog.Loggers or duplicate underlying io.Writers be added.
func (logger *Logger) AddLogger(newLogger *log.Logger) error {
	for _, l := range logger.logs {
		if l == newLogger {
			return errors.New("duplicate logger added")
		}
		if l.Writer() == newLogger.Writer() {
			return errors.New("duplicate os.Writer added")
		}
	}
	logger.logs = append(logger.logs, newLogger)
	return nil
}

// Debug writes an unformatted information message to the selected
// szLog.Logger if debug level messages are enabled.
func (logger *Logger) Debug(msg ...any) {
	if logger.IsDebug {
		logger.output(debugLabel, fmt.Sprint(msg...))
	}
}

// Debugf writes a formatted information message to the selected szLog.Logger
// if debug level messages are enabled.
func (logger *Logger) Debugf(msgFmt string, msgArgs ...any) {
	if logger.IsDebug {
		logger.output(debugLabel, fmt.Sprintf(msgFmt, msgArgs...))
	}
}

// Info writes an unformatted information message to the selected szLog.Logger
// if iformation level messages are enabled.
func (logger *Logger) Info(msg ...any) {
	if logger.IsInfo {
		logger.output(infoLabel, fmt.Sprint(msg...))
	}
}

// Infof writes a formatted information message to the selected szLog.Logger
// if information level messages are enabled.
func (logger *Logger) Infof(msgFmt string, msgArgs ...any) {
	if logger.IsInfo {
		logger.output(infoLabel, fmt.Sprintf(msgFmt, msgArgs...))
	}
}

// Warn writes an unformatted error message to the selected szLog.Logger if
// warning level messages are enabled.
func (logger *Logger) Warn(msg ...any) {
	if logger.IsWarn {
		logger.output(warnLabel, fmt.Sprint(msg...))
	}
}

// Warnf writes a formatted error message to the selected szLog.Logger if
// warning level messages are enabled.
func (logger *Logger) Warnf(msgFmt string, msgArgs ...any) {
	if logger.IsWarn {
		logger.output(warnLabel, fmt.Sprintf(msgFmt, msgArgs...))
	}
}

// Error logs an unformatted error message to the selected szLog.Logger.
// Error level is always enabled.
func (logger *Logger) Error(msg ...any) {
	logger.output(errorLabel, fmt.Sprint(msg...))
}

// Errorf logs an unformatted error message to the selected szLog.Logger.
// Error level is always enabled.
func (logger *Logger) Errorf(msgFmt string, msgArgs ...any) {
	logger.output(errorLabel, fmt.Sprintf(msgFmt, msgArgs...))
}

// Close is a convenience function calling Close() on the provided io.Closer
// and logging an unformatted error message to the selected szLog.Logger
// should an error occur.  Good for use in defered close operations.
func (logger *Logger) Close(closable io.Closer, args ...any) {
	err := closable.Close()

	if err != nil {
		msg := ""
		if len(args) > 0 {
			msg += " " + fmt.Sprint(args...)
		}
		logger.Error("Close", msg, " caused: ", err)
	}
}

// Closef is a convenience function calling close on the provided io.Closer
// and logging a formatted error message to the selected szLog.Logger
// should an error occur.  Good for use in defered close operations.
func (logger *Logger) Closef(
	closable io.Closer, fmtMsg string, fmtArgs ...any,
) {
	err := closable.Close()

	if err != nil {
		msg := fmt.Sprintf(fmtMsg, fmtArgs...)
		if len(msg) > 0 {
			msg = " " + msg
		}
		logger.Error("Close", msg, " caused: ", err)
	}
}

// Define the package level standard objects.

// Define the standard szLog.logger object.
var std *Logger = New(ErrorLevel, log.Default())

// Mirror the std.IsXXXX booleans for symetric access.
var (
	IsWarn  bool
	IsInfo  bool
	IsDebug bool
)

// SetLevel sets the logging level for the standard szLog.Logger.
func SetLevel(newLevel Level) Level {
	origLevel := std.SetLevel(newLevel)
	IsWarn = std.IsWarn
	IsInfo = std.IsInfo
	IsDebug = std.IsDebug
	return origLevel
}

// AddWriter wraps the provided io.Writer in a new log.Logger and adds it
// to the logs output by the standard szLog.Logger.
func AddWriter(newWriter io.Writer, prefix string, flags int) error {
	return std.AddWriter(newWriter, prefix, flags)
}

// AddLogger adds the provided log.Logger to the logs output by the
// provided szLog.Logger.  Checks are made and an error is returned should
// duplicate szLog.Loggers or duplicate underlying io.Writers be added.
func AddLogger(newLogger *log.Logger) error {
	return std.AddLogger(newLogger)
}

// Debug writes an unformatted information message to the standard
// szLog.Logger if debug level messages are enabled.
func Debug(msg ...any) {
	if IsDebug {
		std.output(debugLabel, fmt.Sprint(msg...))
	}
}

// Debugf writes a formatted information message to the standard szLog.Logger
// if debug level messages are enabled.
func Debugf(msgFmt string, msgArgs ...any) {
	if IsDebug {
		std.output(debugLabel, fmt.Sprintf(msgFmt, msgArgs...))
	}
}

// Info writes an unformatted information message to the standard szLog.Logger
// if iformation level messages are enabled.
func Info(msg ...any) {
	if IsInfo {
		std.output(infoLabel, fmt.Sprint(msg...))
	}
}

// Infof writes a formatted information message to the standard szLog.Logger
// if warning level messages are enabled.
func Infof(msgFmt string, msgArgs ...any) {
	if IsInfo {
		std.output(infoLabel, fmt.Sprintf(msgFmt, msgArgs...))
	}
}

// Warn writes an unformatted warning message to the standard szLog.Logger if
// warning level messages are enabled.
func Warn(msg ...any) {
	if IsWarn {
		std.output(warnLabel, fmt.Sprint(msg...))
	}
}

// Warnf writes a formatted warning message to the standard szLog.Logger if
// warning level messages are enabled.
func Warnf(msgFmt string, msgArgs ...any) {
	if IsWarn {
		std.output(warnLabel, fmt.Sprintf(msgFmt, msgArgs...))
	}
}

// Error writes an unformatted error message to the standard szLog.Logger.
// Error level is always enabled.
func Error(msg ...any) {
	std.output(errorLabel, fmt.Sprint(msg...))
}

// Errorf writes a formatted error message to the standard szLog.Logger.
// Error level is always enabled.
func Errorf(msgFmt string, msgArgs ...any) {
	std.output(errorLabel, fmt.Sprintf(msgFmt, msgArgs...))
}

// Close is a convenience function calling Close() on the provided io.Closer
// and logging an unformatted error message to the standard logger should an
// error occur.  Good for use in defered close operations.
func Close(closable io.Closer, args ...any) {
	std.Close(closable, args...)
}

// Closef is a convenience function calling close on the provided io.Closer
// and logging a formatted error message to the standard logger should an
// error occur.  Good for use in defered close operations.
func Closef(closable io.Closer, fmtMsg string, fmtArgs ...any) {
	std.Closef(closable, fmtMsg, fmtArgs...)
}

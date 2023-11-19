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

//nolint:goCritic // ok
package szLog

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/dancsecs/szTest"
)

func runDefaultLogTest() {
	SetLevel(ErrorLevel)
	Debug("1-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	Info("1-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	Warn("1-", "WE SHOULD ", "NOT SEE", " THIS WARNING MESSAGE")
	Error("1-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	SetLevel(WarnLevel)
	Debug("2-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	Info("2-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	Warn("2-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	Error("2-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	SetLevel(InfoLevel)
	Debug("3-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	Info("3-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	Warn("3-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	Error("3-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	SetLevel(DebugLevel)
	Debug("4-", "WE SHOULD ", "SEE", " THIS DEBUG MESSAGE")
	Info("4-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	Warn("4-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	Error("4-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
}

func runDefaultLogTestIf() {
	SetLevel(ErrorLevel)
	if IsDebug {
		Debug("1-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	}
	if IsInfo {
		Info("1-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	}
	if IsWarn {
		Warn("1-", "WE SHOULD ", "NOT SEE", " THIS WARNING MESSAGE")
	}
	Error("1-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	SetLevel(WarnLevel)
	if IsDebug {
		Debug("2-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	}
	if IsInfo {
		Info("2-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	}
	if IsWarn {
		Warn("2-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	}
	Error("2-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	SetLevel(InfoLevel)
	if IsDebug {
		Debug("3-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	}
	if IsInfo {
		Info("3-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	}
	if IsWarn {
		Warn("3-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	}
	Error("3-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	SetLevel(DebugLevel)
	if IsDebug {
		Debug("4-", "WE SHOULD ", "SEE", " THIS DEBUG MESSAGE")
	}
	if IsInfo {
		Info("4-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	}
	if IsWarn {
		Warn("4-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	}
	Error("4-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
}

func runDefaultLogTestf() {
	SetLevel(ErrorLevel)
	Debugf("1-WE SHOULD %s THIS DEBUG MESSAGE", "NOT SEE")
	Infof("1-WE SHOULD %s THIS INFO MESSAGE", "NOT SEE")
	Warnf("1-WE SHOULD %s THIS WARNING MESSAGE", "NOT SEE")
	Errorf("1-WE SHOULD %s THIS ERROR MESSAGE", "SEE")

	SetLevel(WarnLevel)
	Debugf("2-WE SHOULD %s THIS DEBUG MESSAGE", "NOT SEE")
	Infof("2-WE SHOULD %s THIS INFO MESSAGE", "NOT SEE")
	Warnf("2-WE SHOULD %s THIS WARNING MESSAGE", "SEE")
	Errorf("2-WE SHOULD %s THIS ERROR MESSAGE", "SEE")

	SetLevel(InfoLevel)
	Debugf("3-WE SHOULD %s THIS DEBUG MESSAGE", "NOT SEE")
	Infof("3-WE SHOULD %s THIS INFO MESSAGE", "SEE")
	Warnf("3-WE SHOULD %s THIS WARNING MESSAGE", "SEE")
	Errorf("3-WE SHOULD %s THIS ERROR MESSAGE", "SEE")

	SetLevel(DebugLevel)
	Debugf("4-WE SHOULD %s THIS DEBUG MESSAGE", "SEE")
	Infof("4-WE SHOULD %s THIS INFO MESSAGE", "SEE")
	Warnf("4-WE SHOULD %s THIS WARNING MESSAGE", "SEE")
	Errorf("4-WE SHOULD %s THIS ERROR MESSAGE", "SEE")
}

func runLogTest(logger *Logger) {
	logger.SetLevel(ErrorLevel)
	logger.Debug("1-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	logger.Info("1-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	logger.Warn("1-", "WE SHOULD ", "NOT SEE", " THIS WARNING MESSAGE")
	logger.Error("1-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
	logger.SetLevel(WarnLevel)
	logger.Debug("2-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	logger.Info("2-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	logger.Warn("2-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	logger.Error("2-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
	logger.SetLevel(InfoLevel)
	logger.Debug("3-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	logger.Info("3-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	logger.Warn("3-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	logger.Error("3-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
	logger.SetLevel(DebugLevel)
	logger.Debug("4-", "WE SHOULD ", "SEE", " THIS DEBUG MESSAGE")
	logger.Info("4-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	logger.Warn("4-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	logger.Error("4-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
}

func runLogTestIf(logger *Logger) {
	logger.SetLevel(ErrorLevel)
	if logger.IsDebug {
		logger.Debug("1-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	}
	if logger.IsInfo {
		logger.Info("1-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	}
	if logger.IsWarn {
		logger.Warn("1-", "WE SHOULD ", "NOT SEE", " THIS WARNING MESSAGE")
	}
	logger.Error("1-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	logger.SetLevel(WarnLevel)
	if logger.IsDebug {
		logger.Debug("2-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	}
	if logger.IsInfo {
		logger.Info("2-", "WE SHOULD ", "NOT SEE", " THIS INFO MESSAGE")
	}
	if logger.IsWarn {
		logger.Warn("2-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	}
	logger.Error("2-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	logger.SetLevel(InfoLevel)
	if logger.IsDebug {
		logger.Debug("3-", "WE SHOULD ", "NOT SEE", " THIS DEBUG MESSAGE")
	}
	if logger.IsInfo {
		logger.Info("3-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	}
	if logger.IsWarn {
		logger.Warn("3-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	}
	logger.Error("3-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")

	logger.SetLevel(DebugLevel)
	if logger.IsDebug {
		logger.Debug("4-", "WE SHOULD ", "SEE", " THIS DEBUG MESSAGE")
	}
	if logger.IsInfo {
		logger.Info("4-", "WE SHOULD ", "SEE", " THIS INFO MESSAGE")
	}
	if logger.IsWarn {
		logger.Warn("4-", "WE SHOULD ", "SEE", " THIS WARNING MESSAGE")
	}
	logger.Error("4-", "WE SHOULD ", "SEE", " THIS ERROR MESSAGE")
}

func runLogTestf(logger *Logger) {
	logger.SetLevel(ErrorLevel)
	logger.Debugf("1-WE SHOULD %s THIS DEBUG MESSAGE", "NOT SEE")
	logger.Infof("1-WE SHOULD %s THIS INFO MESSAGE", "NOT SEE")
	logger.Warnf("1-WE SHOULD %s THIS WARNING MESSAGE", "NOT SEE")
	logger.Errorf("1-WE SHOULD %s THIS ERROR MESSAGE", "SEE")

	logger.SetLevel(WarnLevel)
	logger.Debugf("2-WE SHOULD %s THIS DEBUG MESSAGE", "NOT SEE")
	logger.Infof("2-WE SHOULD %s THIS INFO MESSAGE", "NOT SEE")
	logger.Warnf("2-WE SHOULD %s THIS WARNING MESSAGE", "SEE")
	logger.Errorf("2-WE SHOULD %s THIS ERROR MESSAGE", "SEE")

	logger.SetLevel(InfoLevel)
	logger.Debugf("3-WE SHOULD %s THIS DEBUG MESSAGE", "NOT SEE")
	logger.Infof("3-WE SHOULD %s THIS INFO MESSAGE", "SEE")
	logger.Warnf("3-WE SHOULD %s THIS WARNING MESSAGE", "SEE")
	logger.Errorf("3-WE SHOULD %s THIS ERROR MESSAGE", "SEE")

	logger.SetLevel(DebugLevel)
	logger.Debugf("4-WE SHOULD %s THIS DEBUG MESSAGE", "SEE")
	logger.Infof("4-WE SHOULD %s THIS INFO MESSAGE", "SEE")
	logger.Warnf("4-WE SHOULD %s THIS WARNING MESSAGE", "SEE")
	logger.Errorf("4-WE SHOULD %s THIS ERROR MESSAGE", "SEE")
}

const expected = "" +
	`E: 1-WE SHOULD SEE THIS ERROR MESSAGE` + "\n" +
	`W: 2-WE SHOULD SEE THIS WARNING MESSAGE` + "\n" +
	`E: 2-WE SHOULD SEE THIS ERROR MESSAGE` + "\n" +
	`I: 3-WE SHOULD SEE THIS INFO MESSAGE` + "\n" +
	`W: 3-WE SHOULD SEE THIS WARNING MESSAGE` + "\n" +
	`E: 3-WE SHOULD SEE THIS ERROR MESSAGE` + "\n" +
	`D: 4-WE SHOULD SEE THIS DEBUG MESSAGE` + "\n" +
	`I: 4-WE SHOULD SEE THIS INFO MESSAGE` + "\n" +
	`W: 4-WE SHOULD SEE THIS WARNING MESSAGE` + "\n" +
	`E: 4-WE SHOULD SEE THIS ERROR MESSAGE` + "\n" +
	""

func Test_SzLog_Unformatted(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	runLogTest(New(DebugLevel, log.Default()))

	chk.Log(expected)
}

func Test_SzLog_Unformatted_If(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	runLogTestIf(New(DebugLevel, log.Default()))

	chk.Log(expected)
}

func Test_SzLog_Formatted(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	runLogTestf(New(DebugLevel, log.Default()))

	chk.Log(expected)
}

func Test_SzLog_Default_Unformatted(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	runDefaultLogTest()

	chk.Log(expected)
}

func Test_SzLog_Default_UnformattedIf(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	runDefaultLogTestIf()

	chk.Log(expected)
}

func Test_SzLog_Default_Formatted(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	runDefaultLogTestf()

	chk.Log(expected)
}

func Test_SzLog_ToAdditionalFile(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	fDir := chk.CreateTmpDir()
	fPath := filepath.Join(fDir, "testlogfile.log")
	f, err := os.Create(fPath)
	if chk.NoErr(err) {
		defer Close(f)
	}

	chk.NoErr(AddWriter(f, "", 0))
	runDefaultLogTest()
	b, err := ioutil.ReadFile(fPath)
	chk.NoErr(err)

	chk.Str(
		string(b),
		expected,
	)
	chk.Log(expected)
}

func Test_SzLog_ToAdditionalLogger(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	fDir := chk.CreateTmpDir()
	fPath := filepath.Join(fDir, "testlogfile.log")
	f, err := os.Create(fPath)
	if chk.NoErr(err) {
		defer Close(f)
	}
	logger := log.New(f, "", 0)

	chk.NoErr(AddLogger(logger))
	runDefaultLogTest()
	b, err := ioutil.ReadFile(fPath)
	chk.NoErr(err)

	chk.Str(
		string(b),
		expected,
	)
	chk.Log(expected)
}

func Test_SzLog_DuplicateWriters(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	fDir := chk.CreateTmpDir()
	fPath := filepath.Join(fDir, "testlogfile.log")
	f, err := os.Create(fPath)
	if chk.NoErr(err) {
		defer Close(f)
	}
	logger1 := log.New(f, "", 0)
	logger2 := log.New(f, "", 0)

	chk.NoErr(AddLogger(logger1))
	chk.Err(AddLogger(logger2), "duplicate os.Writer added")

	chk.Log()
}

func Test_SzLog_DuplicateLoggers(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	fDir := chk.CreateTmpDir()
	fPath := filepath.Join(fDir, "testlogfile.log")
	f, err := os.Create(fPath)
	if chk.NoErr(err) {
		defer Close(f)
	}
	logger := log.New(f, "", 0)

	chk.NoErr(AddLogger(logger))
	chk.Err(AddLogger(logger), "duplicate logger added")

	chk.Log()
}

func TestSzLog_MultilineLogMessage(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	Errorf("This is line 1 of message\nAnd this is line 2 of message")
	chk.Log("" +
		"E: This is line 1 of message\n" +
		"+  And this is line 2 of message\n" +
		"")
}

func TestSzLog_CloseFile(t *testing.T) {
	chk := szTest.CaptureLog(t)
	defer chk.Release()

	_ = chk.CreateTmpDir()
	tstFile := chk.CreateTmpFile([]byte("file contents"))

	f, err := os.Open(tstFile)
	chk.NoErr(err)

	Close(f)
	Close(f)
	Close(f, "msg1", " msg2")
	Closef(f, "msg1 msg2")
	Closef(f, "msg1 %s", "msg2")

	chk.AddSub(`{{tstFile}}`, tstFile)

	chk.Log("" +
		`E: Close caused: close {{tstFile}}: file already closed` + "\n" +
		`E: Close msg1 msg2 caused: close {{tstFile}}: file already closed` + "\n" +
		`E: Close msg1 msg2 caused: close {{tstFile}}: file already closed` + "\n" +
		`E: Close msg1 msg2 caused: close {{tstFile}}: file already closed` + "\n" +
		"")
}

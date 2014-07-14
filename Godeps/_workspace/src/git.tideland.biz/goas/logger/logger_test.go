// Tideland Go Application Support - Logger - Unit Tests
//
// Copyright (C) 2012-2013 Frank Mueller / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package logger_test

//--------------------
// IMPORTS
//--------------------

import (
	"log"
	"os"
	"testing"

	"git.tideland.biz/goas/logger"
	"git.tideland.biz/gots/asserts"
)

//--------------------
// TESTS
//--------------------

// Test log level.
func TestLevel(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	logger.SetLevel(logger.LevelDebug)
	assert.Equal(logger.Level(), logger.LevelDebug, "Level debug.")
	logger.SetLevel(logger.LevelCritical)
	assert.Equal(logger.Level(), logger.LevelCritical, "Level critical.")
	logger.SetLevel(logger.LevelDebug)
	assert.Equal(logger.Level(), logger.LevelDebug, "Level debug.")
}

// Test debugging.
func TestDebug(t *testing.T) {
	logger.Debugf("Hello, I'm debugging %v!", "here")
	logger.SetLevel(logger.LevelError)
	logger.Debugf("Should not be shown!")
}

// Test log at all levels.
func TestAllLevels(t *testing.T) {
	logger.SetLevel(logger.LevelDebug)

	logger.Debugf("Debug.")
	logger.Infof("Info.")
	logger.Warningf("Warning.")
	logger.Errorf("Error.")
	logger.Criticalf("Critical.")
}

// Test logging from level warning and above.
func TestWarningAndAbove(t *testing.T) {
	logger.SetLevel(logger.LevelWarning)

	logger.Debugf("Debug.")
	logger.Infof("Info.")
	logger.Warningf("Warning.")
	logger.Errorf("Error.")
	logger.Criticalf("Critical.")
}

// Test logging with the go logger.
func TestGoLogger(t *testing.T) {
	log.SetOutput(os.Stdout)

	logger.SetLevel(logger.LevelDebug)
	logger.SetLogger(logger.GoLogger{})

	logger.Debugf("Debug.")
	logger.Infof("Info.")
	logger.Warningf("Warning.")
	logger.Errorf("Error.")
	logger.Criticalf("Critical.")
}

// Test logging with an own logger.
func TestOwnLogger(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
	ownLogger := &testLogger{[]string{}}

	logger.SetLevel(logger.LevelDebug)
	logger.SetLogger(ownLogger)

	logger.Debugf("Debug.")
	logger.Infof("Info.")
	logger.Warningf("Warning.")
	logger.Errorf("Error.")
	logger.Criticalf("Critical.")

	assert.Length(ownLogger.logs, 5, "Everything logged.")
}

//--------------------
// LOGGER
//--------------------

type testLogger struct {
	logs []string
}

func (tl *testLogger) Debug(info, msg string) {
	tl.logs = append(tl.logs, info+" "+msg)
}

func (tl *testLogger) Info(info, msg string) {
	tl.logs = append(tl.logs, info+" "+msg)
}
func (tl *testLogger) Warning(info, msg string) {
	tl.logs = append(tl.logs, info+" "+msg)
}
func (tl *testLogger) Error(info, msg string) {
	tl.logs = append(tl.logs, info+" "+msg)
}
func (tl *testLogger) Critical(info, msg string) {
	tl.logs = append(tl.logs, info+" "+msg)
}

// EOF

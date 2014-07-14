// Tideland Go Application Support - Errors - Unit Tests
//
// Copyright (C) 2013 Frank Mueller / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package errors_test

//--------------------
// IMPORTS
//--------------------

import (
	"testing"

	"git.tideland.biz/goas/errors"
	"git.tideland.biz/gots/asserts"
)

//--------------------
// TESTS
//--------------------

// Test creation and checking.
func TestIsError(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
	// Asserts.
	ec := 42
	err := errors.New(ec, "test error %d", 1)
	assert.ErrorMatch(err, `test error 1 \(E042\)`, "error code is added")
	assert.True(errors.IsError(err, ec))
	assert.False(errors.IsError(err, 0))

	err = testError("test error 2")
	assert.ErrorMatch(err, "test error 2")
	assert.False(errors.IsError(err, ec))
	assert.False(errors.IsError(err, 0))
}

// Test correct payload handling.
func TestPayloadOK(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
	// Asserts.
	ec := 42
	err := errors.New(ec, "test error payload ok")
	assert.Nil(errors.SetPayload(err, "foo", "bar"))

	value, errerr := errors.Payload(err, "foo")
	assert.Nil(errerr)
	assert.Equal(value, "bar")

	value, errerr = errors.Payload(err, "baz")
	assert.Nil(value)
	assert.ErrorMatch(errerr, `payload with key "baz" not found \(E002\)`)

	assert.Nil(errors.SetPayload(err, "yadda", nil))
	value, errerr = errors.Payload(err, "yadda")
	assert.Nil(errerr)
	assert.Nil(value)
}

// Test incorrect payload handling.
func TestPayloadFail(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)
	// Asserts.
	err := testError("test error payload fail")
	errerr := errors.SetPayload(err, "foo", "bar")
	assert.ErrorMatch(errerr, `invalid error type: errors_test.testError "test error payload fail" \(E001\)`)

	value, errerr := errors.Payload(err, "foo")
	assert.Nil(value)
	assert.ErrorMatch(errerr, `invalid error type: errors_test.testError "test error payload fail" \(E001\)`)
}

//--------------------
// HELPERS
//--------------------

type testError string

func (e testError) Error() string {
	return string(e)
}

// EOF

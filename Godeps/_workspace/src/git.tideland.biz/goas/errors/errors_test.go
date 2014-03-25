// Tideland Go Application Support - Errors - Unit Tests
//
// Copyright (C) 2013-2014 Frank Mueller / Oldenburg / Germany
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

// Test the validation.
func TestValidation(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	ec := 1
	err := errors.New(ec, "valid")
	fileName, line, lerr := errors.Location(err)

	assert.True(errors.Valid(err))
	assert.Nil(lerr)
	assert.Equal(fileName, "errors_test.go")
	assert.Equal(line, 30)
}

// Test the annotation of errors.
func TestAnnotation(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	ec := 123
	aerr := testError("wrapped")
	err := errors.Annotate(aerr, ec, "annotated")

	assert.ErrorMatch(err, `\[E123\] annotated: wrapped`)
	assert.Equal(errors.Annotated(err), aerr)
	assert.True(errors.IsInvalidTypeError(errors.Annotated(aerr)))
	assert.Length(errors.Stack(err), 2)
}

// Test creation and checking.
func TestIsError(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	ec := 42
	err := errors.New(ec, "test error %d", 1)

	assert.ErrorMatch(err, `\[E042\] test error 1`)
	assert.True(errors.IsError(err, ec))
	assert.False(errors.IsError(err, 0))

	err = testError("test error 2")

	assert.ErrorMatch(err, "test error 2")
	assert.False(errors.IsError(err, ec))
	assert.False(errors.IsError(err, 0))
}

//--------------------
// HELPERS
//--------------------

type testError string

func (e testError) Error() string {
	return string(e)
}

// EOF

// Tideland Go Application Support - Errors - Unit Tests
//
// Copyright (C) 2013-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package errors_test

//--------------------
// IMPORTS
//--------------------

import (
	"testing"

	"github.com/tideland/goas/v3/errors"
	"github.com/tideland/gots/v3/asserts"
)

//--------------------
// TESTS
//--------------------

// Test the validation.
func TestValidation(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	ec := 1
	messages := errors.Messages{ec: "valid"}
	err := errors.New(ec, messages)
	fileName, line, lerr := errors.Location(err)

	assert.True(errors.Valid(err))
	assert.Nil(lerr)
	assert.Equal(fileName, "errors_test.go")
	assert.Equal(line, 31)
}

// Test the annotation of errors.
func TestAnnotation(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	ec := 123
	messages := errors.Messages{ec: "annotated"}
	aerr := testError("wrapped")
	err := errors.Annotate(aerr, ec, messages)

	assert.ErrorMatch(err, `\[E123\] annotated: wrapped`)
	assert.Equal(errors.Annotated(err), aerr)
	assert.True(errors.IsInvalidTypeError(errors.Annotated(aerr)))
	assert.Length(errors.Stack(err), 2)
}

// Test creation and checking.
func TestIsError(t *testing.T) {
	assert := asserts.NewTestingAssertion(t, true)

	ec := 42
	messages := errors.Messages{ec: "test error %d"}
	err := errors.New(ec, messages, 1)

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

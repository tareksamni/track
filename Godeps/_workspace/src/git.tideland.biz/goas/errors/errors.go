// Tideland Go Application Support - Errors
//
// Copyright (C) 2013 Frank Mueller / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package errors

//--------------------
// IMPORTS
//--------------------

import (
	"fmt"
)

//--------------------
// CONSTANTS
//--------------------

const (
	ecNoError = iota
	ecInvalidErrorType
	ecPayloadNotFound

	msgInvalidErrorType = "invalid error type: %T %q"
	msgPayloadNotFound  = "payload with key %q not found"
)

//--------------------
// ERROR
//--------------------

// errorBox encapsulates an error.
type errorBox struct {
	code    int
	msg     string
	payload map[string]interface{}
}

// Error returns the error as string.
func (e *errorBox) Error() string {
	return fmt.Sprintf("%s (E%03d)", e.msg, e.code)
}

// New creates an error with the given code, message and arguments.
func New(code int, format string, args ...interface{}) error {
	return &errorBox{
		code: code,
		msg:  fmt.Sprintf(format, args...),
	}
}

// IsError checks if an error is one created by this
// package and has the passed code
func IsError(err error, code int) bool {
	if e, ok := err.(*errorBox); ok {
		return e.code == code
	}
	return false
}

// IsInvalidTypeError is true if the error signals an invalid
// error type for the payload operations.
func IsInvalidTypeError(err error) bool {
	return IsError(err, ecInvalidErrorType)
}

// IsPayloadNotFoundError is true if the error signals that
// a payload with the given key cannot be found.
func IsPayloadNotFoundError(err error) bool {
	return IsError(err, ecPayloadNotFound)
}

// SetPayload sets a payload inside the error that can
// be analyzed by the receiver.
func SetPayload(err error, key string, value interface{}) error {
	if e, ok := err.(*errorBox); ok {
		if e.payload == nil {
			e.payload = make(map[string]interface{})
		}
		e.payload[key] = value
		return nil
	}
	return New(ecInvalidErrorType, msgInvalidErrorType, err, err)
}

// Payload returns the value of an error payload.
func Payload(err error, key string) (interface{}, error) {
	if e, ok := err.(*errorBox); ok {
		if e.payload != nil {
			if v, ok := e.payload[key]; ok {
				return v, nil
			}
		}
		return nil, New(ecPayloadNotFound, msgPayloadNotFound, key)
	}
	return nil, New(ecInvalidErrorType, msgInvalidErrorType, err, err)
}

// EOF

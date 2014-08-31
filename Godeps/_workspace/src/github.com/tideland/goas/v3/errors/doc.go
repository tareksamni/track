// Tideland Go Application Support - Errors
//
// Copyright (C) 2013-2014 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

// Typical errors in Go are often created using errors.New()
// or fmt.Errorf(). Those errors only contain a string as information.
// When trying to differentiate between errors or to carry helpful
// payload own types are needed.
//
// The errors package allows to easily created formatted errors
// with New() like with the fmt.Errorf function, but also with an
// error code. This easily can be tested with IsError(). So own
// error testing functions like IsMyFooError() bool can simply be
// implemented. If an existing error shall be wrapped the function
// Annotate() allows to pass that error. It can be retrieved using
// Annotated() or Stack() for the full stack.
package errors

// EOF

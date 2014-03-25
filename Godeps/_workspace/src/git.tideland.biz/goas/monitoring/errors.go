// Tideland Go Application Support - Monitoring
//
// Copyright (C) 2009-2013 Frank Mueller / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package monitoring

//--------------------
// IMPORTS
//--------------------

import (
	"git.tideland.biz/goas/errors"
)

//--------------------
// CONSTANTS
//--------------------

const (
	ecNoError = iota
	ecMonitorPanicked
	ecMonitorCannotBeRecovered
	ecMeasuringPointNotExists
	ecStaySetVariableNotExists
	ecDynamicStatusNotExists

	msgMonitorPanicked          = "monitor backend panicked"
	msgMonitorCannotBeRecovered = "monitor cannot be recovered: %v"
	msgMeasuringPointNotExists  = "measuring point %q does not exist"
	msgStaySetVariableNotExists = "stay-set variable %q does not exist"
	msgDynamicStatusNotExists   = "dynamic status %q does not exist"
)

//--------------------
// TESTING
//--------------------

// IsMonitorPanickedError returns true, if the error signals that
// the monitor backend panicked.
func IsMonitorPanickedError(err error) bool {
	return errors.IsError(err, ecMonitorPanicked)
}

// IsMonitorCannotBeRecoveredError returns true, if the error signals that
// the monitor backend has panicked to often and cannot be recovered.
func IsMonitorCannotBeRecoveredError(err error) bool {
	return errors.IsError(err, ecMonitorCannotBeRecovered)
}

// IsMeasuringPointNotExistsError returns true, if the error signals that
// a wanted measuring point cannot be retrieved because it doesn't exists.
func IsMeasuringPointNotExistsError(err error) bool {
	return errors.IsError(err, ecMeasuringPointNotExists)
}

// IsStaySetVariableNotExistsError returns true, if the error signals that
// a wanted stay-set variable cannot be retrieved because it doesn't exists.
func IsStaySetVariableNotExistsError(err error) bool {
	return errors.IsError(err, ecStaySetVariableNotExists)
}

// IsDynamicStatusNotExistsError returns true, if the error signals that
// a wanted dynamic status cannot be retrieved because it doesn't exists.
func IsDynamicStatusNotExistsError(err error) bool {
	return errors.IsError(err, ecDynamicStatusNotExists)
}

// EOF

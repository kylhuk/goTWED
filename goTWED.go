// Package gotwed offers Go bindings for the C library of TWED
package gotwed

/*
#cgo LDFLAGS: -lm
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#include "twed.c"
*/
import "C"

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

/*
TWED compares two time series and calculates the distance between them. If the distance is 0, the two time series are
identical. The bigger the distance, the bigger the difference between the two time series

Input:

	ta []float64 - Time series A, e.g. {5, 1, 3, -8}
	tsa []float64 - Time stamps of time series A, e.g {0, 1, 2, 3}
	tb []float64 - Time series B, e.g. {1, 5, 1, 2}
	tsb []float64 - Time stamps of time series B, e.g {0, 1, 2, 3}
	nu float64 - hyperparameter nu, e.g. 1
	lambda float64 - hyperparameter lambda, e.g. 1
	degree uint32 - power degree for the evaluation of the local distance between samples; degree > 0 required

Output:

	distance float64
*/
func TWED(ta []float64, tsa []float64, tb []float64, tsb []float64, nu float64, lambda float64, degree int32) (distance float64) {

	// Error handling
	if len(ta) != len(tsa) {
		logrus.Error(fmt.Sprintf("Length mismatch: time series ta[] does not have the same length as tsa[]. Length of ta[]: %d Length of tsa[]: %d", len(ta), len(tsa)))
		return -1
	}
	if len(tb) != len(tsb) {
		logrus.Error(fmt.Sprintf("Length mismatch: time series tb[] does not have the same length as tsb[]. Length of tb[]: %d Length of tsb[]: %d", len(tb), len(tsb)))
		return -1
	}
	if degree < 0 {
		logrus.Error(fmt.Sprintf("Parameter degree is negative. This parameter has to be > 0. Value of degree: %d", degree))
		return -1
	}

	dist := C.double(0)  // CTWED writes to this object
	la := C.int(len(ta)) // length of the time series A
	lb := C.int(len(tb)) // length of the time series B

	C.CTWED((*C.double)(&ta[0]), &la, (*C.double)(&tsa[0]), (*C.double)(&tb[0]), &lb, (*C.double)(&tsb[0]), (*C.double)(&nu), (*C.double)(&lambda), (*C.int)(&degree), &dist)

	return float64(dist)
}

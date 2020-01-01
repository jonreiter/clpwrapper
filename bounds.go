package clpwrapper

import (
	"errors"
	"math"

	"github.com/lanl/clp"
	"gonum.org/v1/gonum/mat"
)

// ConvertBounds from a gonum mat into clp bounds
func ConvertBounds(varBounds *mat.Dense) ([]clp.Bounds, error) {
	rows, cols := varBounds.Dims()
	if cols != 2 {
		return nil, errors.New("Bounds matrix must have 2 columns")
	}
	cb := make([]clp.Bounds, rows)
	for i := 0; i < rows; i++ {
		cb[i].Lower = varBounds.At(i, 0)
		cb[i].Upper = varBounds.At(i, 1)
	}
	return cb, nil
}

// BuildAllPositiveBounds returns a vector of 0->math.Inf(1) bounds
func BuildAllPositiveBounds(nv int) []clp.Bounds {
	zeroVec := make([]float64, nv)
	return BuildAboveBounds(zeroVec)
}

// BuildAllNegativeBounds returns a vector of math.Inf(-1)->0 bounds
func BuildAllNegativeBounds(nv int) []clp.Bounds {
	zeroVec := make([]float64, nv)
	return BuildBelowBounds(zeroVec)
}

// BuildEqualityBounds returns a vector of v<->v bounds
func BuildEqualityBounds(boundsVec []float64) []clp.Bounds {
	varBounds := make([]clp.Bounds, len(boundsVec))
	for i, v := range boundsVec {
		varBounds[i].Lower = v
		varBounds[i].Upper = v
	}
	return varBounds
}

// BuildBelowBounds returns a vector of math.Inf(-1)->v bounds
func BuildBelowBounds(boundsVec []float64) []clp.Bounds {
	varBounds := make([]clp.Bounds, len(boundsVec))
	for i, v := range boundsVec {
		varBounds[i].Lower = math.Inf(-1)
		varBounds[i].Upper = v
	}
	return varBounds
}

// BuildAboveBounds returns a vector of math.Inf(-1)->v bounds
func BuildAboveBounds(boundsVec []float64) []clp.Bounds {
	varBounds := make([]clp.Bounds, len(boundsVec))
	for i, v := range boundsVec {
		varBounds[i].Lower = v
		varBounds[i].Upper = math.Inf(1)
	}
	return varBounds
}

// eof

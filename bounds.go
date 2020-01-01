package clpwrapper

import (
	"math"

	"github.com/lanl/clp"
	"gonum.org/v1/gonum/mat"
)

// ConvertBounds from a gonum mat into clp bounds
func ConvertBounds(varBounds *mat.Dense) []clp.Bounds {
	rows, cols := varBounds.Dims()
	if cols != 2 {
		panic("Bounds matrix must have 2 columns")
	}
	cb := make([]clp.Bounds, rows)
	for i := 0; i < rows; i++ {
		cb[i].Lower = varBounds.At(i, 0)
		cb[i].Upper = varBounds.At(i, 1)
	}
	return cb
}

// BuildAllPositiveBounds returns a vector of 0->math.Inf(1) bounds
func BuildAllPositiveBounds(nv int) []clp.Bounds {
	varBoundsMat := mat.NewDense(nv, 2, nil)
	for i := 0; i < nv; i++ {
		varBoundsMat.Set(i, 0, 0)
		varBoundsMat.Set(i, 1, math.Inf(1))
	}
	return ConvertBounds(varBoundsMat)
}

// eof

// Package clpwrapper contains simple wrappers to help with lanl/clp
// once go has an approximately-standard sparse matrix
// package this code can go into the clp itself
package clpwrapper

import (
	"github.com/james-bowman/sparse"
	"github.com/lanl/clp"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// equality tolerance
const tolerance = 1e-5

// GoNumMatrixToCLPPackedMatrix converts a likely-sparse mat.Matrix into a CoinPackedMatrix
func GoNumMatrixToCLPPackedMatrix(matrix mat.Matrix) *clp.PackedMatrix {
	nRows, nCols := matrix.Dims()
	packedMat := clp.NewPackedMatrix()
	packedMat.Reserve(nCols, nRows*nCols, false)

	for c := 0; c < nCols; c++ {
		col := make([]clp.Nonzero, 0)
		for r := 0; r < nRows; r++ {
			thisVal := matrix.At(r, c)
			if !floats.EqualWithinAbs(thisVal, 0.0, tolerance) {
				col = append(col, clp.Nonzero{Index: r, Value: thisVal})
			}
		}
		packedMat.AppendColumn(col)
	}
	return packedMat
}

// CSCToCLPPackedMatrix converts a sparse.CSC into a CoinPackedMatrix
func CSCToCLPPackedMatrix(matrix *sparse.CSC) *clp.PackedMatrix {
	_, nCols := matrix.Dims()
	totalNNZ := matrix.NNZ()
	packedMat := clp.NewPackedMatrix()
	packedMat.Reserve(nCols, totalNNZ, false)

	for c := 0; c < nCols; c++ {
		col := make([]clp.Nonzero, 0)
		matrix.DoColNonZero(c, func(i, j int, v float64) {
			ele := clp.Nonzero{Index: i, Value: v}
			col = append(col, ele)
		})
		packedMat.AppendColumn(col)
	}

	return packedMat
}

// COOToCLPPackedMatrix converts a sparse.COO into a CoinPackedMatrix
func COOToCLPPackedMatrix(matrix *sparse.COO) *clp.PackedMatrix {
	csc := matrix.ToCSC()
	return CSCToCLPPackedMatrix(csc)
}

// eof

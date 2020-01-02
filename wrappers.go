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

// GoNumMatrixToCLPPackedMatrix converts a likely-sparse mat.Matrix into a CoinPackedMatrix
func GoNumMatrixToCLPPackedMatrix(matrix mat.Matrix) *clp.PackedMatrix {
	return GoNumMatrixToCLPPackedMatrixAtTolerance(matrix, 0.0)
}

// GoNumMatrixToCLPPackedMatrixAtTolerance converts a likely-sparse mat.Matrix into a CoinPackedMatrix
// entries with absolute value less than tolerance are dropped
func GoNumMatrixToCLPPackedMatrixAtTolerance(matrix mat.Matrix, tolerance float64) *clp.PackedMatrix {
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

// LoadSparseProblem loads from sparse constraint matrices
func LoadSparseProblem(simp *clp.Simplex, C []float64,
	varBounds, ubBounds []clp.Bounds, eqBoundsVec []float64,
	AEQ, AUB *sparse.COO) {

	nRowsEQ, nColsEQ := AEQ.Dims()
	nRowsUB, _ := AUB.Dims()

	AEQcsc := AEQ.ToCSC()
	AUBcsc := AUB.ToCSC()

	// merge together the A_eq and A_ub matrices
	mergedCOO := sparse.NewCOO(nRowsEQ+nRowsUB, nColsEQ, nil, nil, nil)
	for c := 0; c < len(C); c++ {
		AEQcsc.DoColNonZero(c, func(i, j int, v float64) {
			mergedCOO.Set(i, j, v)
		})
		AUBcsc.DoColNonZero(c, func(i, j int, v float64) {
			mergedCOO.Set(i+nRowsEQ, j, v)
		})
	}
	mergedCSC := mergedCOO.ToCSC()

	// merge the bounds
	eqBounds := BuildEqualityBounds(eqBoundsVec)
	fullBounds := append(eqBounds, ubBounds...)

	// Load the problem into the model.
	cpm := CSCToCLPPackedMatrix(mergedCSC)
	simp.LoadProblem(cpm, varBounds, C, fullBounds, nil)
}

// eof

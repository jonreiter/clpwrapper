package clpwrapper_test

import (
	"math"
	"testing"

	"github.com/jonreiter/clpwrapper"
	"gonum.org/v1/gonum/mat"

	"github.com/james-bowman/sparse"
	"github.com/lanl/clp"
)

func TestLoadSparseProblem(t *testing.T) {
	C := []float64{1, 1, 1}

	eqbv := []float64{6}
	aeq := sparse.NewCOO(1, 3, nil, nil, nil)
	aeq.Set(0, 0, 1)

	varBounds := []clp.Bounds{
		{Lower: 1, Upper: 6},
		{Lower: 1, Upper: 6},
		{Lower: 1, Upper: 6},
	}

	ubBounds := []clp.Bounds{
		{Lower: math.Inf(-1), Upper: -1},
		{Lower: math.Inf(-1), Upper: -1},
		{Lower: math.Inf(-1), Upper: -1},
	}

	aub := sparse.NewCOO(3, 3, nil, nil, nil)
	aubm := mat.NewDense(3, 3, nil)
	aub.Set(0, 0, -1)
	aub.Set(0, 1, 1)
	aub.Set(1, 1, -1)
	aub.Set(1, 2, 1)
	aub.Set(2, 0, 1)
	aub.Set(2, 1, -2)
	aub.Set(2, 2, 1)

	aubm.Set(0, 0, -1)
	aubm.Set(0, 1, 1)
	aubm.Set(1, 1, -1)
	aubm.Set(1, 2, 1)
	aubm.Set(2, 0, 1)
	aubm.Set(2, 1, -2)
	aubm.Set(2, 2, 1)

	simp := clp.NewSimplex()
	clpwrapper.LoadSparseProblem(simp, C, varBounds, ubBounds, eqbv, aeq, aub)
	status := simp.Primal(clp.NoValuesPass, clp.NoStartFinishOptions)
	if status != clp.Optimal {
		t.Error("got back non-optimal")
	}
	res := simp.PrimalColumnSolution()
	if res[0] != 6 || res[1] != 4 || res[2] != 1 {
		t.Error("bad solution")
	}

	clpm1 := clpwrapper.COOToCLPPackedMatrix(aub)
	clpm2 := clpwrapper.GoNumMatrixToCLPPackedMatrix(aubm)
	dd1 := clpm1.DenseData()
	dd2 := clpm2.DenseData()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if dd1[i][j] != dd2[i][j] {
				t.Error("dense data mismatch")
			}
		}
	}

}

// TestZeros passes in a constraint matrix that is almost-entirely 0s.
// This requires explicitly setting the size of the PackedMatrix.
func TestZeros(t *testing.T) {
	C := []float64{1}

	ubBounds := []clp.Bounds{
		{Lower: 0, Upper: 0},
		{Lower: 0, Upper: 0},
		{Lower: 0, Upper: 0},
		{Lower: 0, Upper: 0},
	}

	simp := clp.NewSimplex()

	packedMat := clp.NewPackedMatrix()
	packedMat.Reserve(1, 1, false)
	packedMat.AppendColumn([]clp.Nonzero{{Index: 0, Value: 1}})
	packedMat.SetDimensions(4, 1)

	simp.LoadProblem(packedMat, nil, C, ubBounds, nil)

}

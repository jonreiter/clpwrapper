package clpwrapper_test

import (
	"math"
	"testing"

	"github.com/jonreiter/clpwrapper"
	"github.com/lanl/clp"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/mat"
)

func TestConvertBounds(t *testing.T) {
	src := rand.NewSource(1)
	rnd := rand.New(src)
	r := 10
	b := mat.NewDense(r, 2, nil)
	for i := 0; i < r; i++ {
		b.Set(i, 0, rnd.Float64())
		b.Set(i, 1, b.At(i, 0)+rnd.Float64())
	}
	clpbs, err := clpwrapper.ConvertBounds(b)
	if err != nil {
		t.Error("convert error")
	}
	if !boundsMatch(t, b, clpbs) {
		t.Error("convertbounds mismatch")
	}

	b2 := mat.NewDense(r, 3, nil)
	_, err = clpwrapper.ConvertBounds(b2)
	if err == nil {
		t.Error("didnt get expected error")
	}
}

func TestBuildBounds(t *testing.T) {
	src := rand.NewSource(1)
	rnd := rand.New(src)
	r := 4 + rnd.Intn(100)
	b := mat.NewDense(r, 2, nil)
	for i := 0; i < r; i++ {
		b.Set(i, 0, 0)
		b.Set(i, 1, math.Inf(1))
	}
	clpbs := clpwrapper.BuildAllPositiveBounds(r)
	if !boundsMatch(t, b, clpbs) {
		t.Error("buildallpositive mismatch")
	}
	for i := 0; i < r; i++ {
		b.Set(i, 0, math.Inf(-1))
		b.Set(i, 1, 0)
	}
	clpbs = clpwrapper.BuildAllNegativeBounds(r)
	if !boundsMatch(t, b, clpbs) {
		t.Error("buildallnegative mismatch")
	}
	bvec := make([]float64, r)
	for i := 0; i < r; i++ {
		bvec[i] = rnd.Float64()
		b.Set(i, 0, bvec[i])
		b.Set(i, 1, bvec[i])
	}
	clpbs = clpwrapper.BuildEqualityBounds(bvec)
	if !boundsMatch(t, b, clpbs) {
		t.Error("buildequality mismatch")
	}
	for i := 0; i < r; i++ {
		bvec[i] = rnd.Float64()
		b.Set(i, 0, math.Inf(-1))
		b.Set(i, 1, bvec[i])
	}
	clpbs = clpwrapper.BuildBelowBounds(bvec)
	if !boundsMatch(t, b, clpbs) {
		t.Error("buildbelow mismatch")
	}
	for i := 0; i < r; i++ {
		bvec[i] = rnd.Float64()
		b.Set(i, 0, bvec[i])
		b.Set(i, 1, math.Inf(1))
	}
	clpbs = clpwrapper.BuildAboveBounds(bvec)
	if !boundsMatch(t, b, clpbs) {
		t.Error("buildabove mismatch")
	}
}

func boundsMatch(t *testing.T, b mat.Matrix, clpbs []clp.Bounds) bool {
	t.Helper()
	for i, clpb := range clpbs {
		if b.At(i, 0) != clpb.Lower {
			t.Error("lower bound mismatch")
			return false
		}
		if b.At(i, 1) != clpb.Upper {
			t.Error("upper bound mismatch")
			return false
		}
	}
	return true
}

// gofst.go
package gofst

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L/usr/local/lib -lfst
// #include "gofst.h"
import "C"

//Fst structy
type Fst struct {
	fst C.CFst
}

//New create a new Fst object
func New() Fst {
	var ret Fst
	ret.fst = C.FstInit()
	return ret
}

//Free free the fst object
func (f Fst) Free() {
	C.FstFree(f.fst)
}

//AddState add a new state for fst
func (f Fst) AddState() {
	C.AddState(f.fst)
}

//SetStart set a new id for start state
func (f Fst) SetStart(state C.int) {
	C.SetStart(f.fst, state)
}

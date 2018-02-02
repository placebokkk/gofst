// gofst.go
package gofst

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L/usr/local/lib -lfst
// #include "gofst.h"
import "C"

//Fst structy
type Fst struct {
	fst      C.CFst
	symTable C.CSymbolTable
}

//FstNew create a new Fst object
func FstNew() Fst {
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

//Compose compose two fst to a new fst
func (f Fst) Compose(f2 Fst) Fst {
	ofst := FstNew()
	C.Compose(f.fst, f2.fst, ofst.fst)
	return ofst
}

//ArcSortInput sort output arc
func (f Fst) ArcSortInput() {
	C.ArcSortInput(f.fst)
}

//ArcSortOuput sort input arc
func (f Fst) ArcSortOuput() {
	C.ArcSortOutput(f.fst)
}

//ArcSortOuput sort input arc
func (f Fst) Write(filename string) {
	C.FstWrite(f.fst, (C.CString)(filename))
}

//ArcSortOuput sort input arc
func FstRead(filename string) Fst {
	var ret Fst
	ret.fst = C.FstRead((C.CString)(filename))
	return ret
}

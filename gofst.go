// gofst.go
package gofst

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L/usr/local/lib -lfst
// #include "gofst.h"
import "C"

//Fst structy
type Fst struct {
	cfst       C.CFst
	cisyms C.CSymbolTable
	cosyms C.CSymbolTable
}

//FstNew create a new Fst object
func FstInit() Fst {
	var ret Fst
	ret.cfst = C.FstInit()
	return ret
}

//Free free the fst object
func (f Fst) Free() {
	C.FstFree(f.cfst)
}

//AddState add a new state for fst
func (f Fst) AddState() {
	C.FstAddState(f.cfst)
}

//SetStart set a new id for start state
func (f Fst) SetStart(state C.int) {
	C.FstSetStart(f.cfst, state)
}

//fst.AddArc(1, StdArc(3, 3, 2.5, 2));

//AddState add a new state for fst
func (f Fst) AddArc(state_id int, arc Arc) {
	C.FstAddArc(f.cfst, C.int(state_id), arc.carc)
}

//OPERATION
//In pyfst, fst object also carries the isyms and osysm.
//We want to use the same design.
//TODO support isyms and osyms
func (f Fst) Determinize() Fst {
	ofst := FstInit()
	C.FstDeterminize(f.cfst, ofst.cfst)
	return ofst
}

//Compose compose two fst to a new fst
func (f Fst) Compose(f2 Fst) Fst {
	ofst := FstInit()
	C.FstCompose(f.cfst, f2.cfst, ofst.cfst)
	return ofst
}


//ArcSortInput sort output arc
func (f Fst) ArcSortInput() {
	C.FstArcSortInput(f.cfst)
}

//ArcSortOuput sort input arc
func (f Fst) ArcSortOuput() {
	C.FstArcSortOutput(f.cfst)
}

//I/O

//Write write FST to file
func (f Fst) Write(filename string) {
	C.FstWrite(f.cfst, (C.CString)(filename))
}

//FstRead read FST from file
func FstRead(filename string) Fst {
	var ret Fst
	ret.cfst = C.FstRead((C.CString)(filename))
	return ret
}

// Iterator

//C++ openfst usage
//   for (StateIterator<StdFst> siter(fst);
//        !siter.Done();
//        siter.Next()) {
//     StateId s = siter.Value();
//   }

//Go version usage
//   for siter := StateIteratorInit(fst);
//			!siter.Done();
//			siter.Next() {
//		fmt.Println(siter.Value())
//	 }

// State Iterator

type StateIterator struct {
	csiter C.CStateIterator
}

func StateIteratorInit(fst Fst) StateIterator {
	var siter StateIterator
	siter.csiter = C.StateIteratorInit(fst.cfst)
	return siter
}

func (siter StateIterator) Next() {
	C.StateIteratorNext(siter.csiter)
}

func (siter StateIterator) Value() int {
	return int(C.StateIteratorValue(siter.csiter))
}

func (siter StateIterator) Done() bool {
	return int(C.StateIteratorDone(siter.csiter)) > 0
}

// Arc Iterator

type ArcIterator struct {
	caiter C.CArcIterator
}

type Arc struct {
	carc C.CArc
}

func (arc Arc) GetILabel() int {
	return int(C.ArcGetILabel(arc.carc))
}

func (arc Arc) GetOLabel() int {
	return int(C.ArcGetOLabel(arc.carc))
}

func (arc Arc) GetWeight() float64 {
	return float64(C.ArcGetWeight(arc.carc))
}

func (arc Arc) GetNextState() int {
	return int(C.ArcGetNextState(arc.carc))
}

// using StdArc = ArcTpl<TropicalWeight>;
// using Weight = W;
// using Label = int;
// using StateId = int;
// using TropicalWeight = TropicalWeightTpl<float>;
// ArcTpl(Label ilabel, Label olabel, Weight weight, StateId nextstate)

//StdArc(3, 3, 2.5, 2);

func ArcInit(ilabel int, olabel int, weight float64, state_id int) Arc {
	var ret Arc
	ret.carc = C.ArcInit(C.int(ilabel), C.int(olabel), C.float(weight), C.int(state_id))
	return ret
}

func ArcIteratorInit(fst Fst, state_id int) ArcIterator {
	var aiter ArcIterator
	aiter.caiter = C.ArcIteratorInit(fst.cfst, C.int(state_id))
	return aiter
}

func (aiter ArcIterator) Next() {
	C.ArcIteratorNext(aiter.caiter)
}

//how does this box operation cost? should just use C.CArc?
func (aiter ArcIterator) Value() Arc {
	var ret Arc
	ret.carc = C.ArcIteratorValue(aiter.caiter)
	return ret
}

func (aiter ArcIterator) Done() bool {
	return int(C.ArcIteratorDone(aiter.caiter)) > 0
}

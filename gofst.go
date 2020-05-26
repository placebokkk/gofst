//Package gofst
package gofst

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L/usr/local/lib -lfst
// #include "gofst.h"
import "C"

//Fst structy
type Fst struct {
	cfst  C.CFst
	isyms SymbolTable
	osyms SymbolTable
}

type SymbolTable struct {
	csyms C.CSymbolTable
}

//FstNew create a new Fst object
func FstInit() Fst {
	var ret Fst
	ret.cfst = C.FstInit()
	//ret.isyms = isyms
	//ret.isyms = isyms
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
func (f Fst) SetStart(state int) {
	C.FstSetStart(f.cfst, C.int(state))
}

//GetStart set a new id for start state
func (f Fst) GetStart() int {
	return int(C.FstGetStart(f.cfst))
}

//IsFinal set a new id for start state
func (f Fst) IsFinal(state int) bool {
	return C.FstIsFinal(f.cfst, C.int(state)) > 0
}

//SetStart set a new id for start state
func (f Fst) SetFinal(state int, weight float64) {
	C.FstSetFinal(f.cfst, C.int(state), C.float(weight))
}

//fst.AddArc(1, StdArc(3, 3, 2.5, 2));

//AddState add a new state for fst
func (f Fst) AddArcRaw(stateId int, arc Arc) {
	C.FstAddArc(f.cfst, C.int(stateId), arc.carc)
}

//memory leak here, looks like problem in SymbolTable funcs.
func (f Fst) AddArc(src int, tgt int, isym string, osym string, weight float64) {
	var ilabel, olabel int
	if f.isyms.HasSymbol(isym) {
		ilabel = f.isyms.FindKey(isym)
	} else {
		ilabel = f.isyms.AddSymbol(isym)
	}
	if f.osyms.HasSymbol(osym) {
		olabel = f.osyms.FindKey(osym)
	} else {
		olabel = f.osyms.AddSymbol(osym)
	}
	arc := ArcInit(int(ilabel), int(olabel), weight, tgt)
	C.FstAddArc(f.cfst, C.int(src), arc.carc)
}

//memory leak free variant of adding arc
func (f Fst) AddArcBySymbolKey(src int, tgt int, isym int, osym int, weight float64) {
	arc := ArcInit(isym, osym, weight, tgt)
	C.FstAddArc(f.cfst, C.int(src), arc.carc)
}

func (f Fst) Copy() Fst {
	var ret Fst
	ret.cfst = C.FstCopy(f.cfst)
	return ret
}

//OPERATION
//In pyfst, fst object also carries the isyms and osysm.
//We want to use the same design.
//TODO bugs! stringWeight not equal! What is stringWeight
//and how is that used here
func (f Fst) Determinize() Fst {
	ofst := FstInit()
	C.FstDeterminize(f.cfst, ofst.cfst)
	return ofst
}

//Compose compose two fst to a new fst
//TODO support isyms and osyms verification
func (f Fst) Compose(f2 Fst) Fst {
	ofst := FstInit()
	C.FstCompose(f.cfst, f2.cfst, ofst.cfst)
	return ofst
}

//RmEpsilon
func (f Fst) RmEpsilon() Fst {
	C.FstRmEpsilon(f.cfst)
	return f
}

//Invert
func (f Fst) Invert() Fst {
	C.FstInvert(f.cfst)
	return f
}

//Minimize
func (f Fst) Minimize() Fst {
	C.FstMinimize(f.cfst)
	return f
}

//ShortestPath
func (f Fst) ShortestPath(n int) Fst {
	ofst := FstInit()
	C.FstShortestPath(f.cfst, ofst.cfst, C.int(n))
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

//SetInputSymbols set FST input SymbolTable
func (f *Fst) SetInputSymbols(st SymbolTable) {
	f.isyms = st
}

//SetOutputSymbols set FST output SymbolTable
func (f *Fst) SetOutputSymbols(st SymbolTable) {
	f.osyms = st
}

//GetInputSymbols set FST output SymbolTable
func (f *Fst) InputSymbols() SymbolTable {
	return f.isyms
}

//GetOutputSymbols set FST output SymbolTable
func (f *Fst) OutputSymbols() SymbolTable {
	return f.osyms
}

//FstRead read FST from file
func FstRead(filename string) Fst {
	var ret Fst
	ret.cfst = C.FstRead((C.CString)(filename))
	return ret
}

//SymbolTable
func SymbolTableInit() SymbolTable {
	var ret SymbolTable
	ret.csyms = C.SymbolTableInit()
	//ret.isyms = isyms
	//ret.isyms = isyms
	return ret
}

func (st SymbolTable) FindKey(symbol string) int {
	return int(C.SymbolTableFindKey(st.csyms, C.CString(symbol)))
}

//memory leak here, what causes it? maybe csymbol?
func (st SymbolTable) FindSymbol(key int) string {
	//defer C.FreeString(csymbol)
	return C.GoString(C.SymbolTableFindSymbol(st.csyms, C.int(key)))

}

func (st SymbolTable) HasKey(key int) bool {
	return C.SymbolTableHasKey(st.csyms, C.int(key)) > 0
}

func (st SymbolTable) HasSymbol(symbol string) bool {
	return C.SymbolTableHasSymbol(st.csyms, C.CString(symbol)) > 0
}

func (st SymbolTable) AddSymbol(symbol string) int {
	return int(C.SymbolTableAddSymbol(st.csyms, C.CString(symbol)))
}

func (st SymbolTable) AddSymbolKey(symbol string, key int) int {
	return int(C.SymbolTableAddSymbolKey(st.csyms, C.CString(symbol), C.int(key)))
}

func SymbolTableReadText(filename string) SymbolTable {
	var ret SymbolTable
	ret.csyms = C.SymbolTableReadText((C.CString)(filename))
	return ret
}

func SymbolTableRead(filename string) SymbolTable {
	var ret SymbolTable
	ret.csyms = C.SymbolTableReadBinary((C.CString)(filename))
	return ret
}

//Write write FST to file
func (st SymbolTable) Write(filename string) {
	C.SymbolTableWrite(st.csyms, (C.CString)(filename))
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

//return [[arc, arc,...], [arc,arc,..]..]
//a [][]Arc
func (fst Fst) Paths() [][]Arc {
	return _visit(fst, fst.GetStart(), []Arc{})
}

func _visit(fst Fst, state_id int, prefix_path []Arc) [][]Arc {
	var ret [][]Arc = make([][]Arc, 0)
	if fst.IsFinal(state_id) {
		ret = append(ret, prefix_path)
	}
	for aiter := ArcIteratorInit(fst, state_id); !aiter.Done(); aiter.Next() {
		arc := aiter.Value()
		paths := _visit(fst, arc.GetNextState(), append(prefix_path, arc))
		for _, path := range paths {
			ret = append(ret, path)
		}

	}
	return ret
}

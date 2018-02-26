// gofst_test.go
package gofst

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	fst := FstInit()
	fst.AddState()
	fst.SetStart(0)
	fst.Free()
}

func TestCompose(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	model := FstRead("ex01/lexicon_opt.fst")
	input.ArcSortOuput()
	model.ArcSortInput()
	result := input.Compose(model)
	result.Write("composed.fst")
}

func TestDeterminize(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	result := input.Determinize()
	result.Write("determinied.fst")
}

func TestRmEpsilon(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	result := input.RmEpsilon()
	result.Write("rmepsiloned.fst")
}

func TestInvert(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	result := input.Invert()
	result.Write("inverted.fst")
}

func TestMinimize(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	result := input.Minimize()
	result.Write("minimized.fst")
}

func TestStateIterator(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	fmt.Println("start state iterate")
	for siter := StateIteratorInit(input); !siter.Done(); siter.Next() {
		fmt.Println(siter.Value())
	}
}

func TestArcIterator(t *testing.T) {
	input := FstRead("ex01/Marsman_t.fst")
	fmt.Println("start state iterate")
	for siter := StateIteratorInit(input); !siter.Done(); siter.Next() {
		state := siter.Value()
		fmt.Println(state)
		for aiter := ArcIteratorInit(input, state); !aiter.Done(); aiter.Next() {
			arc := aiter.Value()
			fmt.Println(arc.GetILabel(), arc.GetOLabel(), arc.GetWeight(), arc.GetNextState())
		}
	}
}

func TestSymbolTableReadText(t *testing.T) {
	syms := SymbolTableReadText("ex01/ascii.syms")
	for i := 33; i < 50; i++ {
		symbol := syms.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms.FindKey(symbol))
	}
}

func TestSymbolTableRead(t *testing.T) {
	syms := SymbolTableRead("data/lexicon/isyms.fst")
	for i := 0; i <= 8; i++ {
		symbol := syms.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms.FindKey(symbol))
	}
}

func TestSymbolTableWrite(t *testing.T) {
	syms := SymbolTableInit()

	syms.AddSymbolKey("开始", 9)
	fmt.Println("add 开始")
	fmt.Println(9)

	i := syms.AddSymbol("上")
	fmt.Println("add 上")
	fmt.Println(i)

	i = syms.AddSymbol("海")
	fmt.Println("add 海")
	fmt.Println(i)

	syms.Write("data/test.syms.fst")

	syms2 := SymbolTableRead("data/test.syms.fst")
	for i := 9; i <= 11; i++ {
		symbol := syms2.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms2.FindKey(symbol))
	}
}

func TestSymbolTableHasSymbol(t *testing.T) {
	syms := SymbolTableInit()

	syms.AddSymbolKey("开始", 9)
	fmt.Println("add 开始")
	fmt.Println(9)

	i := syms.AddSymbol("上")
	fmt.Println("add 上")
	fmt.Println(i)

	i = syms.AddSymbol("海")
	fmt.Println("add 海")
	fmt.Println(i)

	syms.Write("data/test.syms.fst")

	syms2 := SymbolTableRead("data/test.syms.fst")
	for i := 9; i <= 11; i++ {
		symbol := syms2.FindSymbol(i)
		fmt.Println(syms2.HasKey(i))
		fmt.Println(symbol)
		fmt.Println(syms2.HasSymbol(symbol))
	}
}

func TestSymbolTableSetSymbolTable(t *testing.T) {
	fst := FstInit()
	isyms := SymbolTableInit()
	fst.SetInputSymbols(isyms)

	syms1 := fst.isyms
	syms1.Write("dummy.syms.fst")
	// syms1.AddSymbolKey("input", 0)
	// fmt.Println("add input")
	// fmt.Println(9)

	// i := syms1.AddSymbol("上")
	// fmt.Println("add 上")
	// fmt.Println(i)

	// i = syms1.AddSymbol("交")
	// fmt.Println("add 交")
	// fmt.Println(i)

	// fmt.Println("isyms")
	// for i := 0; i <= 2; i++ {
	// 	symbol := syms1.FindSymbol(i)
	// 	fmt.Println(symbol)
	// 	fmt.Println(syms1.FindKey(symbol))
	// }

}

func TestFstAddArc(t *testing.T) {
	fst := FstInit()
	fst.SetInputSymbols(SymbolTableInit())
	fst.SetOutputSymbols(SymbolTableInit())

	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.AddState()

	fst.AddArc(0, 1, "上", "上", 0.1)
	fst.AddArc(1, 2, "海", "海", 0.4)
	fst.AddArc(1, 3, "交", "学", 0.3)
	fst.SetStart(0)
	fst.Write("addarc.fst")

	fmt.Println("isyms")
	syms1 := fst.isyms

	for i := 0; i <= 2; i++ {
		symbol := syms1.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms1.FindKey(symbol))
	}

	fmt.Println("osyms")
	syms2 := fst.osyms
	for i := 0; i <= 2; i++ {
		symbol := syms2.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms2.FindKey(symbol))
	}

	fmt.Println("read fst from file")
	fst = FstRead("addarc.fst")

	fmt.Println("start state iterate")
	for siter := StateIteratorInit(fst); !siter.Done(); siter.Next() {
		state := siter.Value()
		fmt.Println(state)
		for aiter := ArcIteratorInit(fst, state); !aiter.Done(); aiter.Next() {
			arc := aiter.Value()
			fmt.Println(arc.GetILabel(), arc.GetOLabel(), arc.GetWeight(), arc.GetNextState())
		}
	}

}

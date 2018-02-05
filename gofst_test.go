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

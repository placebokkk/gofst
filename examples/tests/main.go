package main

import (
	"fmt"

	"github.com/placebokkk/gofst"
)

func main() {
	TestBasic()
	TestIsFinal()
	TestFstPaths()
}

func TestBasic() {
	fst := gofst.FstInit()
	fst.AddState()
	fst.SetStart(0)
	fst.Free()
}

func TestIsFinal() {
	fst := gofst.FstInit()
	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.SetStart(0)
	fst.SetFinal(2, 0.1)
	fmt.Println(fst.IsFinal(0))
	fmt.Println(fst.IsFinal(1))
	fmt.Println(fst.IsFinal(2))
}

func TestCompose() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	model := gofst.FstRead("ex01/lexicon_opt.fst")
	input.ArcSortOuput()
	model.ArcSortInput()
	result := input.Compose(model)
	result.Write("composed.fst")
}

func TestDeterminize() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	result := input.Determinize()
	result.Write("determinied.fst")
}

func TestRmEpsilon() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	result := input.RmEpsilon()
	result.Write("rmepsiloned.fst")
}

func TestInvert() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	result := input.Invert()
	result.Write("inverted.fst")
}

func TestMinimize() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	result := input.Minimize()
	result.Write("minimized.fst")
}

func TestStateIterator() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	fmt.Println("start state iterate")
	for siter := gofst.StateIteratorInit(input); !siter.Done(); siter.Next() {
		fmt.Println(siter.Value())
	}
}

func TestArcIterator() {
	input := gofst.FstRead("ex01/Marsman_t.fst")
	fmt.Println("start state iterate")
	for siter := gofst.StateIteratorInit(input); !siter.Done(); siter.Next() {
		state := siter.Value()
		fmt.Println(state)
		for aiter := gofst.ArcIteratorInit(input, state); !aiter.Done(); aiter.Next() {
			arc := aiter.Value()
			fmt.Println(arc.GetILabel(), arc.GetOLabel(), arc.GetWeight(), arc.GetNextState())
		}
	}
}

func TestSymbolTableReadText() {
	syms := gofst.SymbolTableReadText("ex01/ascii.syms")
	for i := 33; i < 50; i++ {
		symbol := syms.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms.FindKey(symbol))
	}
}

func TestSymbolTableRead() {
	syms := gofst.SymbolTableRead("data/lexicon/isyms.fst")
	for i := 0; i <= 8; i++ {
		symbol := syms.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms.FindKey(symbol))
	}
}

func TestSymbolTableWrite() {
	syms := gofst.SymbolTableInit()

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

	syms2 := gofst.SymbolTableRead("data/test.syms.fst")
	for i := 9; i <= 11; i++ {
		symbol := syms2.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms2.FindKey(symbol))
	}
}

func TestSymbolTableHasSymbol() {
	syms := gofst.SymbolTableInit()

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

	syms2 := gofst.SymbolTableRead("data/test.syms.fst")
	for i := 9; i <= 11; i++ {
		symbol := syms2.FindSymbol(i)
		fmt.Println(syms2.HasKey(i))
		fmt.Println(symbol)
		fmt.Println(syms2.HasSymbol(symbol))
	}
}

func TestSymbolTableSetSymbolTable() {
	fst := gofst.FstInit()
	isyms := gofst.SymbolTableInit()
	fst.SetInputSymbols(isyms)

	syms1 := fst.InputSymbols()
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

func TestFstAddArc() {
	fst := gofst.FstInit()
	fst.SetInputSymbols(gofst.SymbolTableInit())
	fst.SetOutputSymbols(gofst.SymbolTableInit())

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
	syms1 := fst.InputSymbols()

	for i := 0; i <= 2; i++ {
		symbol := syms1.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms1.FindKey(symbol))
	}

	fmt.Println("osyms")
	syms2 := fst.OutputSymbols()
	for i := 0; i <= 2; i++ {
		symbol := syms2.FindSymbol(i)
		fmt.Println(symbol)
		fmt.Println(syms2.FindKey(symbol))
	}

	fmt.Println("read fst from file")
	fst = gofst.FstRead("addarc.fst")

	fmt.Println("start state iterate")
	for siter := gofst.StateIteratorInit(fst); !siter.Done(); siter.Next() {
		state := siter.Value()
		fmt.Println(state)
		for aiter := gofst.ArcIteratorInit(fst, state); !aiter.Done(); aiter.Next() {
			arc := aiter.Value()
			fmt.Println(arc.GetILabel(), arc.GetOLabel(), arc.GetWeight(), arc.GetNextState())
		}
	}

}

func TestFstPaths() {
	fst := gofst.FstInit()
	fst.SetInputSymbols(gofst.SymbolTableInit())
	fst.SetOutputSymbols(gofst.SymbolTableInit())

	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.AddState()
	fst.AddState()

	fst.AddArc(0, 1, "上", "上", 0.1)
	fst.AddArc(1, 4, "海", "海", 0.2)
	fst.AddArc(4, 5, "大", "大", 0.3)
	fst.AddArc(0, 3, "北", "北", 0.1)
	fst.AddArc(3, 4, "京", "京", 0.05)
	fst.AddArc(5, 6, "学", "学", 0.6)
	fst.AddArc(6, 7, "的", "的", 0.7)

	fst.SetStart(0)
	fst.SetFinal(6, 0.1)
	fst.SetFinal(7, 0.1)

	fst.Write("addarc.fst")
	paths := fst.Paths()

	for idx, path := range paths {
		fmt.Println(idx)
		for _, arc := range path {
			fmt.Print(arc.GetILabel())
			fmt.Print("-")
			fmt.Print(arc.GetOLabel())
			fmt.Print("-")
			fmt.Print(arc.GetNextState())
			fmt.Print("-")
			fmt.Println(arc.GetWeight())
		}
	}

	//test shortestpath
	best_path := fst.ShortestPath(1)
	paths = best_path.Paths()
	for idx, path := range paths {
		fmt.Println(idx)
		for _, arc := range path {
			fmt.Print(arc.GetILabel())
			fmt.Print("-")
			fmt.Print(arc.GetOLabel())
			fmt.Print("-")
			fmt.Print(arc.GetNextState())
			fmt.Print("-")
			fmt.Println(arc.GetWeight())
		}
	}

}

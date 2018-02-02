// gofst_test.go
package gofst

import "testing"

func TestBasic(t *testing.T) {
	fst := FstNew()
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
	result.Write("result.fst")
}

// gofst_test.go
package gofst

import "testing"

func TestFst(t *testing.T) {
	fst := New()
	fst.AddState()
	fst.SetStart(0)
	fst.Free()
}

// gofst_test.go
package gofst

import "testing"

func TestFfst(t *testing.T) {
	fst := New()
	fst.AddState()
	fst.SetStart(0)
	fst.Free()
}

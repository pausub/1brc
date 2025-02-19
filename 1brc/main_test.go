package main

import "testing"

func TestEntryPoint(t *testing.T) {
	t.Run("never fail", func(t *testing.T) {
		neverFail(t)
	})
}

func neverFail(t testing.TB) {
	t.Helper()
	if false {
		t.Error("Something gone wrong")
	}
}

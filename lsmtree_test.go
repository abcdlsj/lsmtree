package lsmtree

import "testing"

func TestLSMTree(t *testing.T) {
	lsm := NewLSMTree()

	lsm.Append([]byte("key1"), []byte("value1"))

	if val, ok := lsm.Get([]byte("key1")); !ok || string(val) != "value1" {
		t.Error("failed to get value")
	}
}

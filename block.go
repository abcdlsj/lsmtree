package lsmtree

type Block struct {
	entries []Entry
}

type Entry struct {
	Key   KEY
	Value VALUE
	Len   int
}

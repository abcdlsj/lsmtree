package lsmtree

type Iterator interface {
	Next() (KEY, VALUE, bool)
}

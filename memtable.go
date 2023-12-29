package lsmtree

type MemTable struct {
}

func NewMemTable() *MemTable {
	return &MemTable{}
}

func (l *MemTable) Get(key KEY) (VALUE, bool) {
	panic("implement me")
}

func (l *MemTable) Set(key KEY, value VALUE) {
	panic("implement me")
}

func (l *MemTable) Iterate() Iterator {
	panic("implement me")
}

func (l *MemTable) IterateRange(start, end KEY) Iterator {
	panic("implement me")
}

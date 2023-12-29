package lsmtree

type LSMTree struct {
	memtable *MemTable
	sstables []*SSTable
	wal      *WAL
}

func (l *LSMTree) Get(key KEY) (VALUE, bool) {
	if val, ok := l.memtable.Get(key); ok {
		return val, ok
	}

	for _, sstable := range l.sstables {
		if val, ok := sstable.Get(key); ok {
			return val, ok
		}
	}

	return nil, false
}

func (l *LSMTree) Append(key KEY, value VALUE) {
	l.memtable.Set(key, value)
	l.wal.Append(key, value)

	// TODO: flush

	// TODO: compact
}

func NewLSMTree() *LSMTree {
	return &LSMTree{
		memtable: NewMemTable(),
		sstables: make([]*SSTable, 0),
		wal:      &WAL{},
	}
}

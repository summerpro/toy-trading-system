package database

import "sync"

type MemDb struct {
	memDb map[string][]byte
	mutex sync.Mutex
}

func NewMemDb(initSize int) *MemDb {
	return &MemDb{
		memDb: make(map[string][]byte, initSize),
		mutex: sync.Mutex{},
	}
}

func (m *MemDb) Set(key, value []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.memDb[string(key)] = value
}

func (m *MemDb) Get(key []byte) []byte {
	value, ok := m.memDb[string(key)]
	if ok {
		return value
	}
	return nil
}

func (m *MemDb) Commit() {

}

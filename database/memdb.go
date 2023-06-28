package database

import "sync"

type MemDb struct {
	memDb map[string][]byte
	mutex sync.Mutex
}

func (m *MemDb) Set(key, value []byte) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.memDb[string(key)] = value
}

func (m *MemDb) Get(key []byte) []byte {
	return m.memDb[string(key)]
}

func (m *MemDb) Commit() {

}

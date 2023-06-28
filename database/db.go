package database

type DB interface {
	Set(key, value []byte)
	Get(key []byte) []byte
	Commit()
}

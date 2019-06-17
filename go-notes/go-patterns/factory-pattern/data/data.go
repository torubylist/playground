package data

type Store interface {
	Open(string) (string, error)
}

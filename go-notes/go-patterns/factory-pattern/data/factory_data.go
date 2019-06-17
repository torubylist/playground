package data


type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	MemStorage
	SSDStorage
)

func NewStorage(t StorageType) Store  {
	switch t {
	case DiskStorage:
		return newDiskStorage()
	case MemStorage:
		return newMemStorage()
	case SSDStorage:
		return newSSDStorage()
	}
}

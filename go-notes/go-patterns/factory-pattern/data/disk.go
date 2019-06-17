package data

type disk int

func newDiskStorage() disk {
	return disk(1)
}

func (d disk)Open(s string) (string, error)  {
	return "", nil
}
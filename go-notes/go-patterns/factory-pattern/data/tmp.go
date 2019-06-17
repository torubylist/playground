package data

type ssd int

func newSSDStorage() ssd {
	return ssd(1)
}

func (m ssd)Open(s string) (string, error)  {
	return "", nil
}

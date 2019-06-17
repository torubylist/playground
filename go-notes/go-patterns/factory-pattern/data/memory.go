package data

type memory int

func newMemStorage() memory {
	return memory(1)
}

func (m memory)Open(s string) (string, error)  {
	return "", nil
}

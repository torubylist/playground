package mapch

//map将一个channel映射成另外一个channel， channel的类型可以不同。类似于python的map。

func MapChan(in <- chan interface{}, fn func(interface{}) interface{}) <-chan interface{}  {
	out := make(chan interface{})

	if in == nil {
		close(out)
		return out
	}

	go func() {
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

package fan_in_recursive


func fanInRecursive(in ...<-chan interface{}) <-chan interface{}  {
	switch len(in) {
	case 0:
		return nil
	case 1:
		return in[0]
	case 2:
		return mergeTwo(in[0], in[1])
	default:
		m := len(in)/2
		return mergeTwo(fanInRecursive(in[:m]...), fanInRecursive(in[m:]...))
	}
}

func mergeTwo(a, b <-chan  interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <- b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
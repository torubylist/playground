package skip_while

func skipWhile(done <-chan struct{}, in <- chan interface{}, fn func(interface{}) bool) <-chan interface{} {
	valChan := make(chan interface{})

	go func() {
		for {
			select {
			case <- done:
				return
			case v, ok := <- in:
				if !ok {
					return
				}
				if fn(v) {
					continue
				}
				valChan <- v

			}
		}
	}()
	return valChan
}

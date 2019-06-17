package skip_fn


func skipFn(done <- chan struct{}, in <-chan interface{}, fn func(interface{}) bool) <-chan interface{}  {
	valChan := make(chan interface{})
	go func() {
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				if !fn(v) {
					valChan <- v
				}
			case <- done:
				return
			}
		}
	}()
	return valChan
}

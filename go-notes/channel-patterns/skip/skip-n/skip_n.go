package skip_n


func skipN(done <-chan struct{}, in <-chan interface{}, n int) <-chan interface{}  {
	valChan := make(chan interface{})
	go func() {
		defer close(valChan)
		var i int
		for {
			select {
			case v, ok:= <- in:
				if !ok {
					return
				}
				if i >= n {
					valChan <- v
				}
			case <- done:
				return
			}
			i++
		}
	}()
	return valChan
}

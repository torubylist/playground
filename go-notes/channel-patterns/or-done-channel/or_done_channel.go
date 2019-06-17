package or_done_channel

func orDone(done <-chan struct{}, in <-chan interface{}) <-chan interface{}  {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <- done:
				return
			case v, ok := <- in:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <- done:
					return
				}
			}
		}
	}()
	return valStream
}

package take_while

//takeWhile只挑选开头满足fn的数据。


func takeWhile(done <-chan struct{}, inChan <- chan interface{}, fn func(interface{}) bool) <-chan interface{}  {
	valChan := make(chan interface{})
	go func() {
		for {
			select{
			case v, ok := <- inChan:
				if !ok {
					return
				}
				if !fn(v) {
					return
				}
				valChan	<- v
			case <- done:
				return
			}
		}
	}()
	return valChan
}

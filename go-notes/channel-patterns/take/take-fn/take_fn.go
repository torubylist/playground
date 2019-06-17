package take_fn

//takeFn
//takeFn 只筛选满足fn的数据。类似于python中的filter函数。
func takeFn(done <-chan struct{}, inChan <-chan interface{}, fn func(interface{}) bool) <-chan interface{}  {
	valChan := make(chan interface{})
	go func() {
		for  {
			select {
			case v,ok := <- inChan:
				if !ok {
					return
				}
				if fn(v) {
					valChan <- v
				}
			case <- done:
				return
			}
		}
	}()
	return valChan
}

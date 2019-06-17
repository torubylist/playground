package take_n

//takeN
//takeN 读取开头N个数据。

func takeN(done <-chan struct{}, in <-chan interface{}, n int) <-chan interface{}  {
	valChan := make(chan interface{})
	go func() {
		defer close(valChan)
		for i:=0;i<n;i++ {
			select {
			case <- done:
				return
			case v,ok := <- in:
				if !ok {
					return
				}
				valChan <- v
			}
		}
	}()

	return valChan
}
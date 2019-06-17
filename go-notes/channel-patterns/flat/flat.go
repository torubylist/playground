package flat


/*
平展(flat)操作是一个有趣的操作。

如果输入是一个channel,channel中的数据还是相同类型的channel， 那么flat将返回一个输出channel,输出channel中的数据
是输入的各个channel中的数据。

它与扇入不同，扇入的输入channel在调用的时候就是固定的，并且以数组的方式提供，而flat的输入是一个channel的channel，可以运行时
随时的加入channel。
*/

func orDone(done <-chan struct{}, inChan <-chan interface{} ) <-chan interface{}  {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)

		for {
			select {
			case <- done:
				return
			case v, ok := <- inChan:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func flat(done <-chan struct{}, inChan <-chan <-chan interface{}) <-chan interface{}  {
	valStream := make(chan interface{})

	go func() {
		defer close(valStream)

		for {
			var stream <-chan interface{}
			select {
			case v, ok := <- inChan:
				if !ok {
					return
				}
				stream = v
			case <- done:
				return
			}

			for v := range orDone(done, stream) {
				select {
				case valStream <- v:
				case <- done:
				}
			}
		}
	}()
	return valStream
}
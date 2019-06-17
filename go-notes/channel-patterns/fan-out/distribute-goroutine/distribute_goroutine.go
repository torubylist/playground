package distribute_goroutine

//分布模式将从输入channel中读取的值往输出channel中的其中一个发送。

//roundrobin的方式选择输出channel。

func fanOut(in <-chan interface{}, out []chan interface{})  {
	go func() {
		defer func() {
			for i:=0;i<len(in);i++ {
				close(out[i])
			}
		}()

		var i int
		var n = len(out)
		for v := range in {
			out[i] <- v
			i = (i+1) % n
		}
	}()
}

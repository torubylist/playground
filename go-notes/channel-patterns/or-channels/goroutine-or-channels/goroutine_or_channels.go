package goroutine_or_channels

import "sync"

/*
当你等待多个信号的时候，如果收到任意一个信号， 就执行业务逻辑，忽略其它的还未收到的信号。

举个例子， 我们往提供相同服务的n个节点发送请求，只要任意一个服务节点返回结果，我们就可以执行下面的业务逻辑，
其它n-1的节点的请求可以被取消或者忽略。当n=2的时候，这就是back request模式。 这样可以用资源来换取latency的提升。

需要注意的是，当收到任意一个信号的时候，其它信号都被忽略。如果用channel来实现，只要从任意一个channel中接收到
一个数据，那么所有的channel都可以被关闭了(依照你的实现，但是输出的channel肯定会被关闭)。

or函数可以处理n个channel,它为每个channel启动一个goroutine，只要任意一个goroutine从channel读取到数据，
输出的channel就被关闭掉了。

为了避免并发关闭输出channel的问题，关闭操作只执行一次。

*/

func orChannels(channels ...chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		var once sync.Once

		for _, c := range channels {
			go func(<-chan interface{}) {
				select {
				case <- c:
					once.Do(func() {
						close(out)
					})
				case <- out:
				}

			}(c)
		}

	}()
	return out
}
package fan_in_goroutine

import "sync"

//每个channel起一个goroutine。


func fanIn(in ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(in))
		for _, c := range in {
			go func(c <-chan interface{}) {
				for v := range c {
					out <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

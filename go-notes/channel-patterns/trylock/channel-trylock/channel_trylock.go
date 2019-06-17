package channel_trylock

/*
主要是利用channel边界情况下的阻塞特性实现的。

你还可以将缓存的大小从1改为n,用来处理n个锁(资源)。
*/

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex  {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct {}{}
	return mu
}

func (m *Mutex)Lock()  {
	<- m.ch
}

func (m *Mutex)Unlock()  {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock unlocked mutex")
	}
}

func (m *Mutex)Trylock() bool  {
	select {
	case <- m.ch:
		return true
	default:
	}
	return false
}

func (m *Mutex)IsLocked() bool  {
	return len(m.ch) == 0
}


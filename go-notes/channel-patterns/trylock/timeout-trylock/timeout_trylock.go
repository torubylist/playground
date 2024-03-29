package timeout_trylock

import "time"

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex  {
	m := &Mutex{ch: make(chan struct{}, 1)}
	m.ch <- struct{}{}
	return m
}

func (m *Mutex)Lock()  {
	<-m.ch
}

func (m *Mutex)Unlock()  {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock unlocked mutex")
	}
}

func (m *Mutex)TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <- m.ch:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}

func (m *Mutex)IsLocked() bool  {
	return len(m.ch) == 0
}
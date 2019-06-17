package pool

/*
1. Object pool pattern is useful in cases where object initialization is more expensive than
the object maintenance.
2. If there are spikes in demand as opposed to a steady demand, the maintenance overhead
might overweigh the benefits of an object pool.
3. It has positive effects on performance due to objects being initialized beforehand
*/

type Pool chan *Object

type Object int

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

/*

p := pool.New(2)

select {
case obj := <-p:
    obj.Do(  )

p <- obj
default:
// No more objects left — retry later or fail
return
}
*/


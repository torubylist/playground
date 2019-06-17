package flat

import (
	"testing"
	//"math/rand"
)

func TestFlat(t *testing.T)  {
	in := make(chan <-chan interface{})
	go func() {
		defer close(in)
		for i:=0;i<5;i++ {
			chanForTest := make(chan interface{}, 2)
			chanForTest <- 1
			chanForTest <- 2
			in <- chanForTest
		}
	}()
	done := make(chan struct{})
	var res []interface{}
	for v := range flat(done, in) {
		res = append(res, v)
	}


}

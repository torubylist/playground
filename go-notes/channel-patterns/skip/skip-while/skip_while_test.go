package skip_while

import (
	"testing"
	"time"
	"reflect"
)

func TestTakeWhile(t *testing.T)  {
	in := make(chan interface{})
	go func() {
		defer close(in)
		for i:=0;i<5;i++ {
			in <- i
		}
	}()

	var res []interface{}
	done := make(chan struct{})
	go func() {
		for v := range skipWhile(done, in, func(i interface{}) bool {
			if i.(int) < 3 {
				return true
			}
			return false
		}) {
			res = append(res, v)
		}
	}()
	time.Sleep(1*time.Millisecond)
	close(done)

	if reflect.DeepEqual(res, []interface{}{3,4}) {
		t.Logf("test pass %v\n", res)
	}else {
		t.Errorf("test failed, expect %v, but got %v\n", []interface{}{3,4}, res)
	}

}

package skip_fn


import (
	"testing"
	"time"
	"reflect"
)

func TestTakeFn(t *testing.T)  {
	inChan := make(chan interface{})
	go func() {
		defer close(inChan)
		for i:=0;i<5;i++ {
			inChan <- i
		}
	}()

	var res []interface{}
	done := make(chan struct{})
	go func() {
		for v := range skipFn(done, inChan, func(i interface{}) bool {
			if i.(int) % 2 == 1 {
				return true
			}
			return false
		}) {
			res = append(res, v)
		}
	}()

	time.Sleep(1 * time.Millisecond)
	close(done)

	if reflect.DeepEqual(res, []interface{}{0,2,4}) {
		t.Logf("test pass!\n")
	}else{
		t.Errorf("test failed, expect %v, but got %v", []interface{}{0,2,4}, res)
	}
}

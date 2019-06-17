package take_n

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
		for v := range takeN(done, inChan, 3) {
			res = append(res, v)
		}
	}()

	time.Sleep(1 * time.Millisecond)
	close(done)

	if reflect.DeepEqual(res, []interface{}{0,1,2}) {
		t.Logf("test pass!\n")
	}else{
		t.Errorf("test failed, expect %v, but got %v", []interface{}{1,3}, res)
	}
}

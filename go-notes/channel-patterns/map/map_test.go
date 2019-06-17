package mapch

import (
	"testing"
	"reflect"
)

func double(x interface{}) interface{} {
	y := x.(int) * 2
	return y
}

func TestMapChan(t *testing.T)  {
	in := make(chan interface{})
	go func() {
		defer close(in)
		for i:=0;i<5;i++ {
			in <- i
		}
	}()
	var res []interface{}
	for v := range MapChan(in, double) {
		res = append(res, v)
	}
	if reflect.DeepEqual(res, []interface{}{0,2,4,6,8}) {
		t.Logf("test pass!\n")
	}else {
		t.Errorf("test failed, expect %v, but got %v\n", []int{0,2,4,6,8}, res)
	}
}

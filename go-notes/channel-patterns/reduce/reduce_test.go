package reduce

import (
	"testing"
)

func Max(x, y interface{}) interface{}  {
	if x.(int) < y.(int) {
		return y
	}
	return x
}

func Min(x, y interface{}) interface{}  {
	if x.(int) < y.(int) {
		return x
	}
	return y
}

func Sum(x, y interface{}) interface{}  {
	z := x.(int) + y.(int)
	return z
}

func TestReduceSum(t *testing.T)  {
	chanInt := make(chan interface{})
	go func() {
		defer close(chanInt)
		for i:=0;i<5;i++ {
			chanInt <- i
		}
	}()
	res := Reduce(chanInt, Sum)
	if res != 10 {
		t.Errorf("test failed.expect %v, but got %v\n", 10, res)
	}
	t.Logf("test sum pass, got %v!\n", res)
}


func TestReduceMax(t *testing.T)  {
	chanInt := make(chan interface{})
	go func() {
		defer close(chanInt)
		for i:=0;i<5;i++ {
			chanInt <- i
		}
	}()
	res := Reduce(chanInt, Max)
	if res != 4 {
		t.Errorf("test failed.expect %v, but got %v\n", 4, res)
	}
	t.Logf("test max pass, got %v!\n", res)
}


func TestReduceMin(t *testing.T)  {
	chanInt := make(chan interface{})
	go func() {
		defer close(chanInt)
		for i:=0;i<5;i++ {
			chanInt <- i
		}
	}()
	res := Reduce(chanInt, Min)
	if res != 0 {
		t.Errorf("test failed.expect %v, but got %v\n", 0, res)
	}
	t.Logf("test max pass, got %v!\n", res)
}
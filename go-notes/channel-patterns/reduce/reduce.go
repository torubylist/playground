package reduce

//用`reduce`实现`sum`、`max`、`min`等聚合操作。类似于python的reduce。

func Reduce(in <-chan interface{}, fn func(x, y interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}
	out := <- in
	for v := range in {
		out = fn(out, v)
	}
	return out
}

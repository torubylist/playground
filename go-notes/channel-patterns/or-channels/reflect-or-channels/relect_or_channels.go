package reflect_or_channels

import "reflect"


/*
Go的反射库针对select语句有专门的数据(reflect.SelectCase)和函数(reflect.Select)处理。
所以我们可以利用反射“随机”地从一组可选的channel中接收数据，并关闭输出channel。

这种方式看起来更简洁。
*/
func orChannels(channels ...chan interface{}) <- chan interface{}  {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		reflect.Select(cases)
	}()
	return orDone
}

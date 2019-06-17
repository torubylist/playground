package distribute_reflect

import "reflect"

func fanOut(in <-chan interface{}, out []chan interface{})  {
	go func() {
		defer func() {
			for i:=0;i<len(out);i++ {
				close(out[i])
			}
		}()

		cases := make([]reflect.SelectCase, len(out))
		for i := range cases {
			cases[i].Dir = reflect.SelectSend
			cases[i].Chan = reflect.ValueOf(out[i])
		}

		for v := range in {
			v := v
			for i := range cases {
				cases[i].Send = reflect.ValueOf(v)
			}
			_, _, _ = reflect.Select(cases)
		}
	}()
}

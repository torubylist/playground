package tee_goroutine


/*
扇出模式(FanOut)是将一个输入channel扇出为多个channel。

扇出行为至少可以分为两种：

从输入channel中读取一个数据，发送给每个输入channel，这种模式称之为Tee模式
从输入channel中读取一个数据，在输出channel中选择一个channel发送
本节只介绍第一种情况，第二种情况是dispatch-goroutine里
*/

func fanOut(in <-chan interface{}, out []chan interface{}, async bool)  {
	go func() {
		defer func() {
			for i:=0;i<len(out);i++ {
				close(out[i])
			}
		}()

		for v := range in {
			v := v
			for i:=0;i<len(out);i++ {
				i := i
				if async {
					go func() {
						out[i] <- v
					}()
				}else{
					out[i] <- v
				}
			}
		}
	}()
}
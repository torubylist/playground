package recursive_or_channels

func orChannels(channels ...chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <- orChannels(channels[:m]...):
			case <- orChannels(channels[m:]...):
			}
		}
	}()

	return orDone
}

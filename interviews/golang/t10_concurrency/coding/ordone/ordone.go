package ordone

func orDone(doneChs ...chan int) chan int {
	if len(doneChs) == 0 {
		return nil
	}
	if len(doneChs) == 1 {
		return doneChs[0]
	}

	if len(doneChs) == 2 {
		out := make(chan int)
		go func() {
			defer close(out)
			select {
			case <-doneChs[0]:
				return
			case <-doneChs[1]:
				return
			}
		}()
		return out
	}

	i := len(doneChs) / 2
	return orDone(orDone(doneChs[:i]...), orDone(doneChs[i:]...))
}

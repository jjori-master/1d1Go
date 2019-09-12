package unit33_goroutine

import "sync"

type Rectangle struct {
	width  int
	height int
}

func change(rect *Rectangle, wg *sync.WaitGroup) {
	rect.width += 1
	wg.Done()
}

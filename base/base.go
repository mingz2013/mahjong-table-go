package base

import "sync"

type Processor interface {
	Run()
}

func RunProcessor(wg sync.WaitGroup, p Processor) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		p.Run()

	}()
}

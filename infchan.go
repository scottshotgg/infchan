package infchan

import (
	"fmt"
	"sync"
)

const (
	internalChanSize = 100
)

type (
	InfChan struct {
		mut     *sync.Mutex
		current chan interface{}
		chans   []chan interface{}
	}
)

func New() *InfChan {
	var ch = make(chan interface{}, internalChanSize)

	return &InfChan{
		mut:     &sync.Mutex{},
		current: ch,
		chans: []chan interface{}{
			ch,
		},
	}
}

func (ic *InfChan) Insert(i interface{}) {
	ic.mut.Lock()
	defer ic.mut.Unlock()
	fmt.Println("inserting:", i)

	select {
	case ic.current <- i:

	default:
		var ch = make(chan interface{}, internalChanSize)

		ic.chans = append(ic.chans, ch)
		ch <- i

		ic.current = ch
	}
}

func (ic *InfChan) Remove() <-chan interface{} {
	ic.mut.Lock()
	defer ic.mut.Unlock()

	if len(ic.chans[0]) == 0 {
		if len(ic.chans) > 1 {
			ic.chans = append(ic.chans[1:])
		}
	}

	return ic.chans[0]
}

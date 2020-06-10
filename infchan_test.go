package infchan_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/scottshotgg/infchan"
)

func TestInfChan(t *testing.T) {
	var (
		infChan = infchan.New()
		timer   = time.NewTimer(2 * time.Second)

		i int
	)

	for {
		select {
		case <-timer.C:
			break

		default:
			infChan.Insert(i)
			i++

			time.Sleep(100 * time.Millisecond)
			continue
		}

		break
	}

	time.AfterFunc(2*time.Second, func() {
		os.Exit(0)
	})

	go func() {
		for {
			fmt.Println(infChan.Remove())
			time.Sleep(50 * time.Millisecond)
		}
	}()
}

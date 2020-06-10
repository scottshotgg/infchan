package infchan_test

import (
	"log"
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

	// time.AfterFunc(2*time.Second, func() {
	// 	os.Exit(0)
	// })

	for {
		select {
		case i := <-infChan.Remove():
			log.Println("got:", i)
			time.Sleep(50 * time.Millisecond)
			continue

		default:
		}

		break
	}
}

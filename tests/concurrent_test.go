package tests

import (
	"sync"
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

const goroutines = 100

// runConcurrently runs fn in `goroutines` goroutines and waits for all of them.
func runConcurrently(fn func(i int)) {
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(i int) {
			defer wg.Done()
			fn(i)
		}(i)
	}

	wg.Wait()
}

func TestParse_Concurrent(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	date := unixFromPastDate(time.Minute * 5)

	runConcurrently(func(i int) {
		if i%2 == 0 {
			got, err := ago.Parse(date, ago.OptNoSuffix)
			assertResult(t, "5 minutes", got, err)
			return
		}

		got, err := ago.Parse(date)
		assertResult(t, "5 minutes ago", got, err)
	})
}

func TestParse_WhileConfiguring(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	date := unixFromPastDate(time.Minute * 5)

	runConcurrently(func(i int) {
		if i%2 == 0 {
			ago.Configure(ago.Config{
				Language: ago.LangEn,
				Location: "UTC",
			})
			return
		}

		got, err := ago.Parse(date)
		assertResult(t, "5 minutes ago", got, err)
	})
}

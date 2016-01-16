package main

import (
	"fmt"
	"testing"
	"time"
	."github.com/vycb/gorazortpl/tpl/bench"
	."github.com/vycb/gorazortpl/tpl"
)

func BenchmarkMain(*testing.B) {
	// Parallel benchmark for text/template.Template.Execute on a single object.
	testing.Benchmark(func(b *testing.B) {

		page := &Page{
			Title: "Bob",
			FavoriteColors: []string{"blue", "green", "mauve"},
			Year: time.Now().Year(),
		}

		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.

			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.

				 fmt.Println(Bench(page))
			}
		})

	})
}

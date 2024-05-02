package main

import (
	"time"
)

func setInterval(f func(), t time.Duration) {
	ticker := time.NewTicker(t)
	go func() {
		for range ticker.C {
			f()
		}
	}()
}

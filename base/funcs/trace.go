package funcs

import (
	"log"
	"time"
)

// Trace will record func time cost, need ues by defer
func Trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

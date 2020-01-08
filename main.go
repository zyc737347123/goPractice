// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	start := time.Now()
	funcs.FetchAll(os.Args[1:])
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

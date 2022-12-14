// pump is iam analytics purger that moves the data generated by your iam-authz-server nodes to any back-end.
// It is primarily used to display your analytics data in the iam operating system.
package main

import (
	"go_iam/internal/pump"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	pump.NewApp("iam-pump").Run()
}

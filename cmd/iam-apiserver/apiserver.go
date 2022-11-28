package main

import (
	"go_iam/internal/apiserver"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	// Set the random number seed so that it is random every time.
	rand.Seed(time.Now().UTC().UnixNano())
	// Set the maximum number of CPUs
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	apiserver.NewApp("iam-apiserver").Run()
}

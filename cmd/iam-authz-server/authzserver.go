package main

// authzserver is the server for iam-authz-server.
// It is responsible for serving the ladon authorization request.

import (
	"go_iam/internal/authzserver"
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

	authzserver.NewApp("iam-authz-server").Run()
}

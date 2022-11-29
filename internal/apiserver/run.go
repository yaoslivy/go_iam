package apiserver

import "go_iam/internal/apiserver/config"

// Run runs the specified APIServer. This should never exit.
func Run(cfg *config.Config) error {
	// shutdown，TLS, grpc, mysql
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
	// Gin router, Redis
	// grpc run, http server run, https server run
	return server.PrepareRun().Run()
}

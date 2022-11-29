package apiserver

import (
	"go_iam/internal/apiserver/config"
	"go_iam/internal/apiserver/options"
	"go_iam/pkg/app"
	"go_iam/pkg/log"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.
`

// NewApp creates an App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions() //Get the configuration parameters of the current application.
	application := app.NewApp(
		"IAM API Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()
		// Create a running configuration instance based on a given IAM pump command line or configuration file option.
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		// Run the services according to the parameters.
		return Run(cfg)
	}
}

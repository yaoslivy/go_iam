package store

var client Factory

// Factory defines the iam platform storage interface.
type Factory interface {
	Policies() PolicyStore
	Secrets() SecretStore
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}

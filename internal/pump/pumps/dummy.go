package pumps

import (
	"context"
	"go_iam/pkg/log"
)

// DummyPump  defines a dummy pump with dummy specific options and common options.
type DummyPump struct {
	CommonPumpConfig
}

// New create a dummy pump instance.
func (p *DummyPump) New() Pump {
	newPump := DummyPump{}

	return &newPump
}

// GetName returns the dummy pump name.
func (p *DummyPump) GetName() string {
	return "Dummy Pump"
}

// Init initialize the dummy pump instance.
func (p *DummyPump) Init(conf interface{}) error {
	log.Debug("Dummy Initialized")

	return nil
}

// WriteData write analyzed data to dummy persistent back-end storage.
func (p *DummyPump) WriteData(ctx context.Context, data []interface{}) error {
	log.Infof("Writing %d records", len(data))

	return nil
}

package vmimageagereceiver

import (
	"sync"

	"github.com/open-telemetry/opentelemetry-collector/component"
)

// Receiver is the type that provides Receiver functionaly for the VM image age metrics.
type Receiver struct {
	vmImageAgeCollector *VMImageAgeCollector

	stopOnce  sync.Once
	startOnce sync.Once
}

// Start starts the underlying VM metrics generator.
func (receiver *Receiver) Start(host component.Host) error {
	receiver.startOnce.Do(func() {
		receiver.vmImageAgeCollector.StartCollection()
	})
	return nil
}

// Shutdown stops and cancels the underlying VM metrics generator.
func (receiver *Receiver) Shutdown() error {
	receiver.stopOnce.Do(func() {
		receiver.vmImageAgeCollector.StopCollection()
	})
	return nil
}

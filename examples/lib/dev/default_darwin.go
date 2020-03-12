package dev

import (
	"github.com/kirbo/ble"
	"github.com/kirbo/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}

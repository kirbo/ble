package dev

import (
	"github.com/kirbo/ble"
	"github.com/kirbo/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}

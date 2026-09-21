package victusdbus

import (
	"fmt"
)

type DBus struct {
	client Client
}

func New() (*DBus, error) {
	client, err := newDBusClient()
	if err != nil {
		return nil, fmt.Errorf(
			"unable to create dbus client: %w",
			err,
		)
	}

	return &DBus{
		client: *client,
	}, nil
}

func (d *DBus) SetFansMax() error {
	if err := d.client.setFansMode(FanModeMax); err != nil {
		return fmt.Errorf("unable to set fans `max`: %w", err)
	}

	return nil
}

func (d *DBus) SetFansAuto() error {
	if err := d.client.setFansMode(FanModeAuto); err != nil {
		return fmt.Errorf("unable to set fans `auto`: %w", err)
	}

	return nil
}

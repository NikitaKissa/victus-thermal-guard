package victusdbus

import (
	"github.com/godbus/dbus/v5"
)

const (
	busName       = "dev.radhey.VictusControl1"
	objectPath    = "/dev/radhey/VictusControl1"
	interfaceName = "dev.radhey.VictusControl1"
)

type Client struct {
	object dbus.BusObject
}

func newDBusClient() (*Client, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}

	object := conn.Object(
		busName,
		dbus.ObjectPath(objectPath),
	)

	return &Client{
		object: object,
	}, nil
}

type FanMode string

const (
	FanModeAuto FanMode = "auto"
	FanModeMax  FanMode = "max"
)

func (c *Client) setFansMode(mode FanMode) error {
	return c.object.Call(
		interfaceName+".SetFanMode",
		0,
		string(mode),
	).Err
}

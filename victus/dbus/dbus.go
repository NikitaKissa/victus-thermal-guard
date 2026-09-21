package victusdbus

import (
	"context"
	"fmt"
	"time"

	"github.com/godbus/dbus/v5"
)

const (
	busName       = "dev.radhey.VictusControl1"
	objectPath    = "/dev/radhey/VictusControl1"
	interfaceName = "dev.radhey.VictusControl1"

	callTimeout = time.Second
)

type fanMode string

const (
	fanModeAuto fanMode = "auto"
	fanModeMax  fanMode = "max"
)

type DBus struct {
	conn   *dbus.Conn
	object dbus.BusObject
}

func New() (*DBus, error) {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		return nil, fmt.Errorf("connect to system bus: %w", err)
	}

	return &DBus{
		conn:   conn,
		object: conn.Object(busName, dbus.ObjectPath(objectPath)),
	}, nil
}

func (d *DBus) Close() error {
	return d.conn.Close()
}

func (d *DBus) SetFansMax(ctx context.Context) error {
	return d.setFanMode(ctx, fanModeMax)
}

func (d *DBus) SetFansAuto(ctx context.Context) error {
	return d.setFanMode(ctx, fanModeAuto)
}

func (d *DBus) setFanMode(ctx context.Context, mode fanMode) error {
	ctx, cancel := context.WithTimeout(ctx, callTimeout)
	defer cancel()

	var ok bool

	err := d.object.
		CallWithContext(ctx, interfaceName+".SetFanMode", 0, string(mode)).
		Store(&ok)
	if err != nil {
		return fmt.Errorf("set fans %q: %w", mode, err)
	}

	if !ok {
		return fmt.Errorf("set fans %q: service returned false", mode)
	}

	return nil
}

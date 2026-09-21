package victusdbus

import "log"

type DBus struct {
	// DBus connection
}

func New() (*DBus, error) {
	// initialize DBus
	return &DBus{}, nil
}

func (d *DBus) SetFansMax() error {

	log.Print("max")
	return nil
}

func (d *DBus) SetFansAuto() error {

	log.Print("auto")
	return nil
}

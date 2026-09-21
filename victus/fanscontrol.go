package victus

type Controller interface {
	SetFansMax() error
	SetFansAuto() error
}

type controller struct {
	backend Backend
}

type Backend interface {
	SetFansMax() error
	SetFansAuto() error
}

func NewController(backend Backend) Controller {
	return &controller{
		backend: backend,
	}
}

var isFansMax = false

func (c *controller) SetFansMax() error {
	if isFansMax {
		return nil
	}

	isFansMax = true
	return c.backend.SetFansMax()
}

func (c *controller) SetFansAuto() error {
	if !isFansMax {
		return nil
	}

	isFansMax = false
	return c.backend.SetFansAuto()
}

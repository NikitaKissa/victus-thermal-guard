package victus

import "context"

type Controller interface {
	SetFansMax(ctx context.Context) error
	SetFansAuto(ctx context.Context) error
}

type controller struct {
	backend   Backend
	isFansMax bool
}

type Backend interface {
	SetFansMax(ctx context.Context) error
	SetFansAuto(ctx context.Context) error
}

func NewController(backend Backend) Controller {
	return &controller{
		backend:   backend,
		isFansMax: false,
	}
}

func (c *controller) SetFansMax(ctx context.Context) error {
	if c.isFansMax {
		return nil
	}

	err := c.backend.SetFansMax(ctx)
	if err == nil {
		c.isFansMax = true
	}

	return err
}

func (c *controller) SetFansAuto(ctx context.Context) error {
	if !c.isFansMax {
		return nil
	}

	err := c.backend.SetFansAuto(ctx)
	if err == nil {
		c.isFansMax = false
	}

	return err
}

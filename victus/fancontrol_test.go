package victus

import (
	"context"
	"errors"
	"testing"
)

// fakeBackend считает вызовы и может возвращать заданную ошибку.
// Он также проверяет ctx: отменённый контекст даёт ошибку, как настоящий D-Bus.
type fakeBackend struct {
	maxCalls  int
	autoCalls int
	err       error
}

func (f *fakeBackend) SetFansMax(ctx context.Context) error {
	f.maxCalls++
	if err := ctx.Err(); err != nil {
		return err
	}
	return f.err
}

func (f *fakeBackend) SetFansAuto(ctx context.Context) error {
	f.autoCalls++
	if err := ctx.Err(); err != nil {
		return err
	}
	return f.err
}

var errBackend = errors.New("backend failure")

func TestSetFansMax_IsIdempotent(t *testing.T) {
	fb := &fakeBackend{}
	c := NewController(fb)
	ctx := context.Background()

	for i := 0; i < 3; i++ {
		if err := c.SetFansMax(ctx); err != nil {
			t.Fatalf("call %d: unexpected error: %v", i, err)
		}
	}

	if fb.maxCalls != 1 {
		t.Errorf("backend max calls = %d, want 1", fb.maxCalls)
	}
}

func TestSetFansMax_ErrorKeepsState(t *testing.T) {
	fb := &fakeBackend{err: errBackend}
	c := NewController(fb)
	ctx := context.Background()

	if err := c.SetFansMax(ctx); !errors.Is(err, errBackend) {
		t.Fatalf("error = %v, want %v", err, errBackend)
	}

	// Состояние не должно измениться: следующий вызов снова идёт в бэкенд.
	fb.err = nil
	if err := c.SetFansMax(ctx); err != nil {
		t.Fatalf("retry: unexpected error: %v", err)
	}
	if fb.maxCalls != 2 {
		t.Errorf("backend max calls = %d, want 2", fb.maxCalls)
	}

	// После успеха повторный вызов бэкенд уже не трогает.
	if err := c.SetFansMax(ctx); err != nil {
		t.Fatalf("third call: unexpected error: %v", err)
	}
	if fb.maxCalls != 2 {
		t.Errorf("backend max calls = %d, want 2", fb.maxCalls)
	}
}

func TestSetFansAuto_WithoutMaxDoesNothing(t *testing.T) {
	fb := &fakeBackend{}
	c := NewController(fb)

	if err := c.SetFansAuto(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fb.autoCalls != 0 {
		t.Errorf("backend auto calls = %d, want 0", fb.autoCalls)
	}
}

func TestSetFansAuto_AfterMax(t *testing.T) {
	fb := &fakeBackend{}
	c := NewController(fb)
	ctx := context.Background()

	if err := c.SetFansMax(ctx); err != nil {
		t.Fatalf("max: unexpected error: %v", err)
	}

	for i := 0; i < 2; i++ {
		if err := c.SetFansAuto(ctx); err != nil {
			t.Fatalf("auto call %d: unexpected error: %v", i, err)
		}
	}

	if fb.autoCalls != 1 {
		t.Errorf("backend auto calls = %d, want 1", fb.autoCalls)
	}
}

func TestSetFansAuto_ErrorKeepsMax(t *testing.T) {
	fb := &fakeBackend{}
	c := NewController(fb)
	ctx := context.Background()

	if err := c.SetFansMax(ctx); err != nil {
		t.Fatalf("max: unexpected error: %v", err)
	}

	fb.err = errBackend
	if err := c.SetFansAuto(ctx); !errors.Is(err, errBackend) {
		t.Fatalf("error = %v, want %v", err, errBackend)
	}

	// Контроллер всё ещё считает, что вентиляторы на max, и повторяет попытку.
	fb.err = nil
	if err := c.SetFansAuto(ctx); err != nil {
		t.Fatalf("retry: unexpected error: %v", err)
	}
	if fb.autoCalls != 2 {
		t.Errorf("backend auto calls = %d, want 2", fb.autoCalls)
	}
}

func TestCancelledContextKeepsState(t *testing.T) {
	fb := &fakeBackend{}
	c := NewController(fb)

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	if err := c.SetFansMax(cancelled); !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want %v", err, context.Canceled)
	}

	// Неудачный вызов не должен пометить вентиляторы как max.
	if err := c.SetFansMax(context.Background()); err != nil {
		t.Fatalf("retry: unexpected error: %v", err)
	}
	if fb.maxCalls != 2 {
		t.Errorf("backend max calls = %d, want 2", fb.maxCalls)
	}
}

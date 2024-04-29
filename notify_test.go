package zenity_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dcam0050/zenity"
	"go.uber.org/goleak"
)

func ExampleNotify() {
	zenity.Notify("There are system updates necessary!",
		zenity.Title("Warning"),
		zenity.InfoIcon)
}

func TestNotify_cancel(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := zenity.Notify("text", zenity.Context(ctx))
	if skip, err := skip(err); skip {
		t.Skip("skipping:", err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Error("was not canceled:", err)
	}
}

func TestNotify_examples(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	ExampleNotify()
}

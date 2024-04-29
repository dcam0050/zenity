package zenity_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dcam0050/zenity"
	"go.uber.org/goleak"
)

func ExamplePassword() {
	zenity.Password(zenity.Title("Type your password"))
}

func ExamplePassword_username() {
	zenity.Password(
		zenity.Title("Type your username and password"),
		zenity.Username())
}

func TestPassword_timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/5)
	defer cancel()

	_, _, err := zenity.Password(zenity.Context(ctx))
	if skip, err := skip(err); skip {
		t.Skip("skipping:", err)
	}
	if !os.IsTimeout(err) {
		t.Error("did not timeout:", err)
	}
}

func TestPassword_cancel(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := zenity.Password(zenity.Context(ctx))
	if skip, err := skip(err); skip {
		t.Skip("skipping:", err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Error("was not canceled:", err)
	}
}

func TestPassword_username(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := zenity.Password(zenity.Context(ctx), zenity.Username())
	if skip, err := skip(err); skip {
		t.Skip("skipping:", err)
	}
	if !errors.Is(err, context.Canceled) {
		t.Error("was not canceled:", err)
	}
}

func TestPassword_script(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	tests := []struct {
		name string
		call string
		opts []zenity.Option
		usr  string
		pwd  string
		err  error
	}{
		{name: "Cancel", call: "cancel", err: zenity.ErrCanceled},
		{name: "Password", call: "enter pwd", pwd: "pwd"},
		{name: "User", call: "enter usr and pwd", usr: "usr", pwd: "pwd",
			opts: []zenity.Option{zenity.Username()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zenity.Info(fmt.Sprintf("In the password dialog, %s.", tt.call))
			usr, pwd, err := zenity.Password(tt.opts...)
			if skip, err := skip(err); skip {
				t.Skip("skipping:", err)
			}
			if usr != tt.usr || pwd != tt.pwd || err != tt.err {
				t.Errorf("Password() = %q, %q, %v; want %q, %q, %v", usr, pwd, err, tt.usr, tt.pwd, tt.err)
			}
		})
	}
}

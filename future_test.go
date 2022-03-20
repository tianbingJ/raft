package raft

import (
	"errors"
	"testing"
	"time"
)

func TestDeferFutureSuccess(t *testing.T) {
	var f deferError
	f.init()
	f.respond(nil)
	if err := f.Error(); err != nil {
		t.Fatalf("unexpected error result; got %#v want nil", err)
	}
	if err := f.Error(); err != nil {
		t.Fatalf("unexpected error result; got %#v want nil", err)
	}
}

func TestDeferFutureError(t *testing.T) {
	want := errors.New("x")
	var f deferError
	f.init()
	f.respond(want)
	if got := f.Error(); got != want {
		t.Fatalf("unexpected error result; got %#v want %#v", got, want)
	}
	if got := f.Error(); got != want {
		t.Fatalf("unexpected error result; got %#v want %#v", got, want)
	}
}

func TestDeferFutureConcurrent(t *testing.T) {
	// Food for the race detector.
	want := errors.New("x")
	var f deferError
	f.init()
	go f.respond(want)
	if got := f.Error(); got != want {
		t.Errorf("unexpected error result; got %#v want %#v", got, want)
	}
}

func TestDeferFutureWait(t *testing.T) {
	var f deferError
	f.init()

	interval := 2 * time.Millisecond

	t1 := time.Now()
	go func() {
		time.Sleep(interval)
		f.respond(nil)
	}()

	error := f.Error()
	if error != nil {
		t.Fatal("expect error nil")
	}
	t2 := time.Now()
	if !t2.After(t1.Add(interval)) {
		t.Fatal("expected wait 2 millis")
	}
}

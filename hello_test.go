// this will a hold a test for our Hello function

package main

import "testing"

// t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
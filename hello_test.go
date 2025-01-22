// this will a hold a test for our Hello function

package main

import "testing"

// t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){

	got := Hello("Chris")
	want := "Hello, Chris"

	assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})


}

	func assertCorrectMessage (t testing.TB, got, want string) {
		t.Helper() // By doing this, when it fails, the line number reported will be in our function call rather than inside our test helper.
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
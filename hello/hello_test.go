package main 

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Kara")
  want := "Hello, Kara"

  if got != want {
    t.Errorf("got %q want %q", got , want)
  }
}

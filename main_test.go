package main

import "strings"
import "testing"

func TestGreeting(t *testing.T) {
	g := Greeting()
	if !strings.Contains(g, "Hello") {
		t.Errorf("Greeting is not polite: %#v", g)
	}
}

package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
    result := EvenOrOdd(10)
    if result != "even" {
        t.Errorf("expected: even,1 actual: %s", result)
    }
}

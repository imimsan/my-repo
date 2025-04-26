package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
    result := EvenOrOdd(10)
    if result != "even" {
        t.Errorf("expected:a ev1en, actual: %s", result)
    }
}

package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 15 {
		t.Errorf("Expected length of 15 but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got: %v", d[0])
	}

	if d[len(d)-1] != "Five of Hearts" {
		t.Errorf("Expected Five of hearts but got: %v", d[len(d)-1])
	}
}

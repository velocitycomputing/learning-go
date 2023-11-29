package main

import "testing"

func Test_NewDeckLength(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

func TestNewDeckFirst(t *testing.T) {
	d := newDeck()
	if d[0] != "A ♥" {
		t.Errorf("Expected first card of A ♥, but got %v", d[0])
	}
}
func TestNewDeckLast(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "K ♣" {
		t.Errorf("Expected first card of K ♣, but got %v", d[0])
	}
}

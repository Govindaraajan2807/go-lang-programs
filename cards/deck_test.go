package main

import (
	"os"
	"testing"
)

func TestNewdeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card as Ace of Hearts, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected cards is 16, but we got %v ", len(loadedDeck))
	}

	os.Remove("_decktesting")

}

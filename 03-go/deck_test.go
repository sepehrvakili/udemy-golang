package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("expected %d, got %d", 20, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected %s, got %s", "Ace of Spades", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("expected %s, got %s", "Five of Clubs", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("expected %d, got %d", 20, len(loadedDeck))
	}

	os.Remove("_decktesting")
}

package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(d))
	}

}

func TestSaveToLocalAndLoad(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()
	d.saveToLocal("_decktesting")

	loadDeck := openCardsFromLocal("_decktesting")

	if len(loadDeck) != 52 {
		t.Errorf("Whoopsie!Expectd 52 got %v", len(loadDeck))
	}

	os.Remove("_decktesting")

}

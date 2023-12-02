package succinctly

import (
	"fmt"
	"testing"
)

func TestLib(t *testing.T) {
	fmt.Printf("Testing...\n")
	// Arrange
	dictionary := []string{"I", "wander", "through", "the", "lovely", "woods"}
	Init(dictionary, -1, -1)

	// Act
	for _, word := range dictionary {
		if code, ok := QCode(word); ok {
			if wordFound, wfOk := QWord(code); wfOk {
				if word != wordFound {
					t.Errorf("Expected %s, but got %s", word, wordFound)
				} else {
					fmt.Printf("Word passed: %s code: %s\n", word, code)
				}
			} else {
				t.Errorf("Word not found for word %s  and code %s", word, code)
			}
		} else {
			t.Errorf("Word not found: %s", word)
		}
	}
}

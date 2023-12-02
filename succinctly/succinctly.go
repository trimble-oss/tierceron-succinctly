package succinctly

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var succinctMap = cmap.New[string]()
var succinctQCodeMap = cmap.New[string]()

// Initialize dictionary with all words you will use.
func Init(dictionary []string, min, max int) error {
	sort.Strings(dictionary)

	if min < 0 || min > max {
		min = 2
	}

	if max < 0 || max > 8 {
		max = 8
	}

	for i := min; i < max; i++ {
		for _, word := range dictionary {
			wordHash := sha256.Sum256([]byte(word))
			success := false
			for j := i; i < max; j++ {
				wordCode := fmt.Sprintf("%x", wordHash[:j])
				if !succinctMap.Has(wordCode) {
					succinctMap.Set(wordCode, word)
					succinctQCodeMap.Set(word, wordCode)
					success = true
					break
				}
			}
			if success {
				// next word please
				continue
			} else {
				succinctMap.Clear()
				break
			}
		}
		if len(dictionary) == succinctMap.Count() {
			// Success.
			return nil
		}
	}

	if len(dictionary) == succinctMap.Count() {
		// Success.
		return nil
	}

	return errors.New("non mappable dictionary")
}

// QCode return the wordcode given the word
func QCode(word string) (string, bool) {
	return succinctQCodeMap.Get(word)
}

// QWord return the word given the wordcode
func QWord(wordCode string) (string, bool) {
	return succinctMap.Get(wordCode)
}

package number_dash_text

import (
	"fmt"
	"testing"
)

type unit struct {
	Input string
	Except bool
}

func TestForTestValidity(t *testing.T) {
	words := []unit{{
		Input: "a-1-2-c",
		Except: false,
	}, {
		Input: "a",
		Except: false,
	}, {
		Input: "a-a-c",
		Except: false,
	}, {
		Input: "23-ab-48-caba-56-haha",
		Except: true,
	}}
	for _, word := range words {
		rst := testValidity(word.Input)
		if rst == word.Except {
			fmt.Printf("word %s pass the test", word.Input)
		}
	}
}
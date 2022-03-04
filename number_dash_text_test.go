package number_dash_text

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
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
			fmt.Printf("word %s pass the test\n", word.Input)
		}
	}
}

func TestForAverageNumber(t *testing.T) {
	invalidWords := []string{"a", "a-1", "a-a", "1-a-2"}
	for _, word := range invalidWords {
		_, err := averageNumber(word)
		assert.True(t, errors.Is(err, InvalidInput))
	}

	avg, err := averageNumber("2-a-2-b")
	assert.True(t, err == nil)
	assert.Equal(t, 2.0, avg)

	avg, err = averageNumber("23-ab-48-caba-56-haha")
	assert.True(t, err == nil)
	assert.Equal(t, 42.333333333333336, avg)
}


func TestForWholeStory(t *testing.T) {
	invalidWords := []string{"a", "a-1", "a-a", "1-a-2"}
	for _, word := range invalidWords {
		_, err := wholeStory(word)
		assert.True(t, errors.Is(err, InvalidInput))
	}

	story, err := wholeStory("2-a-2-b")
	assert.True(t, err == nil)
	assert.Equal(t, "a b", story)

	story, err = wholeStory("23-ab-48-caba-56-haha")
	assert.True(t, err == nil)
	assert.Equal(t, "ab caba haha", story)
}

func TestStoryStats(t *testing.T) {
	invalidWords := []string{"a", "a-1", "a-a", "1-a-2"}
	for _, word := range invalidWords {
		_, err := storyStats(word)
		assert.True(t, errors.Is(err, InvalidInput))
	}

	stats, err := storyStats("2-a-2-b")
	assert.True(t, err == nil)
	assert.Equal(t, "a", stats.Shortest)
	assert.Equal(t, "a", stats.Longest)
	assert.Equal(t, 1.0, stats.AverWordLength)
	assert.Equal(t, []string{"a", "b"}, stats.AverWordLengthText)

	stats, err = storyStats("23-ab-48-caba-56-haha")
	assert.True(t, err == nil)
	assert.Equal(t, "ab", stats.Shortest)
	assert.Equal(t, "caba", stats.Longest)
	assert.Equal(t, 3.3333333333333335, stats.AverWordLength)
	assert.Equal(t, []string{"caba", "haha"}, stats.AverWordLengthText)
}
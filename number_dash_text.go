package number_dash_text

import "errors"

type states struct {
	Shortest string
	Longest string
	AverWordLength float32
	AverWordLengthText []string
}

var InvalidInput = errors.New("invalid input")

/**
testValidity check the input string's format is number-text-number-text
the time difficulty should be O(n)
 */
func testValidity(input string) bool {
	return false
}

/**
averageNumber that takes the string, and returns the average number from all the numbers
the time difficulty should be O(n)
 */
func averageNumber(input string) (float32, error) {
	return 0.0, nil
}

/**
wholeStory that takes the string, and returns a text that is composed from all the text words separated by spaces,
e.g. story called for the string 1-hello-2-world should return text: "hello world"
the time difficulty should be O(n)
 */
func wholeStory(input string) (string, error) {
	return "", nil
}

/**
function storyStats that returns four things:
	-the shortest word
	-the longest word
	-the average word length
	-the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
the time difficulty should be O(n)
 */
func storyStats(input string) (*states, error) {
	return &states{}, nil
}

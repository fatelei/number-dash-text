package number_dash_text

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Stats struct {
	Shortest           string
	Longest            string
	AverWordLength     float64
	AverWordLengthText []string
}

type extractResult struct {
	NumAry    []int64
	StringAry []string
}

var InvalidInput = errors.New("invalid input")

// isDigit check word only has number
func isDigit(word string) bool {
	for _, letter := range word {
		if !unicode.IsDigit(letter) {
			return false
		}
	}
	return true
}

// isLetter check word only has letter
func isLetter(word string) bool {
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			return false
		}
	}
	return true
}

func extractInput(input string) (rst *extractResult, err error) {
	if !testValidity(input) {
		return nil, InvalidInput
	}

	strAry := strings.Split(input, "-")
	length := len(strAry)

	rst = &extractResult{
		NumAry: make([]int64, 0, length / 2),
		StringAry:  make([]string, 0, length / 2),
	}

	for i := 0; i < length; i += 2 {
		a := strAry[i]
		b := strAry[i+1]
		intA, _ := strconv.ParseInt(a, 10, 64)
		rst.NumAry = append(rst.NumAry, intA)
		rst.StringAry = append(rst.StringAry, b)
	}
	fmt.Printf("%+v\n", rst)
	return rst, nil
}

/**
testValidity check the input string's format is number-text-number-text
the time difficulty should be O(n)
*/
func testValidity(input string) bool {
	if len(input) == 0 {
		return false
	}

	strAry := strings.Split(input, "-")
	if len(strAry)%2 != 0 {
		return false
	}

	for i := 0; i < len(strAry); i += 2 {
		a := strAry[i]
		b := strAry[i+1]

		if !isDigit(a) || !isLetter(b) {
			return false
		}
	}
	return true
}

/**
averageNumber that takes the string, and returns the average number from all the numbers
the time difficulty should be O(n)
*/
func averageNumber(input string) (float64, error) {
	rst, err := extractInput(input)
	if err != nil {
		return 0.0, err
	}

	var total int64 = 0
	for _, num := range rst.NumAry {
		total += num
	}

	return float64(total) / float64(len(rst.NumAry)), nil
}

/**
wholeStory that takes the string, and returns a text that is composed from all the text words separated by spaces,
e.g. story called for the string 1-hello-2-world should return text: "hello world"
the time difficulty should be O(n)
*/
func wholeStory(input string) (string, error) {
	rst, err := extractInput(input)
	if err != nil {
		return "", err
	}
	return strings.Join(rst.StringAry, " "), nil
}

/**
function storyStats that returns four things:
	-the shortest word
	-the longest word
	-the average word length
	-the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
the time difficulty should be O(n)
*/
func storyStats(input string) (*Stats, error) {
	rst, err := extractInput(input)
	if err != nil {
		return nil, err
	}

	var shortest string
	var longest string
	totalWordLength := 0

	flagLong := -1
	flagShort := 0xfffffffff

	lengthToWordMap := make(map[int][]string)
	totalWords := len(rst.StringAry)

	for _, word := range rst.StringAry {
		wordLen := len(word)
		totalWordLength += wordLen
		if wordLen > flagLong {
			flagLong = wordLen
			longest = word
		}

		if wordLen < flagShort {
			flagShort = wordLen
			shortest = word
		}

		if _, ok := lengthToWordMap[wordLen]; ok {
			lengthToWordMap[wordLen] = append(lengthToWordMap[wordLen], word)
		} else {
			lengthToWordMap[wordLen] = make([]string, 0, totalWords)
			lengthToWordMap[wordLen] = append(lengthToWordMap[wordLen], word)
		}
	}

	stats := Stats{}
	avg := float64(totalWordLength) / float64(totalWords)
	stats.Longest = longest
	stats.Shortest = shortest
	stats.AverWordLength = avg
	stats.AverWordLengthText = make([]string, 0, totalWords)

	roundUp := int(math.Ceil(avg))
	roundDown := int(math.Floor(avg))

	if v, ok := lengthToWordMap[roundUp]; ok {
		stats.AverWordLengthText = append(stats.AverWordLengthText, v...)
	}

	if roundDown != roundUp {
		if v, ok := lengthToWordMap[roundDown]; ok {
			stats.AverWordLengthText = append(stats.AverWordLengthText, v...)
		}
	}
	return &stats, nil
}


/**
wholeStory that takes the string, and returns a text that is composed from all the text words separated by spaces,
e.g. story called for the string 1-hello-2-world should return text: "hello world"
the time difficulty should be O(n)
*/
func generateValidWord(input string) (string, error) {
	rst, err := extractInput(input)
	if err != nil {
		return "", err
	}
	return strings.Join(rst.StringAry, " "), nil
}
package number_dash_text

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomNumberString(n int) string {
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	numAry := rand.Perm(n)
	numStrAry := make([]string, 0, n)
	for _, i := range numAry {
		strI := strconv.Itoa(i)
		numStrAry = append(numStrAry, strI)
	}
	return strings.Join(numStrAry, "")
}

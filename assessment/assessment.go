package assessment

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	allNumReg  = regexp.MustCompile("^[0-9]+$")
	allTextReg = regexp.MustCompile("^[a-zA-Z]+$")
)

/* testValidity:
	Validates A String And Makes Sure It Conforms To The String Spec Of Format:
	num-text-num-text e.g 1-hello-2-world

	Constraints: It Has To Begin With A Number And End With A Word

	Parameters: string

	Returns: Bool

  	Difficulty: Easy
  	Used Time Of Completion: 10 Minutes
*/
func testValidity(text string) bool {
	text = strings.TrimSpace(text)
	var control byte = 0

	// Makes Sure That No Two Dashes Follow Each Other i.e 1--ab-2-cd
	for _, char := range []byte(text) {
		if char == 45 {
			control += 1
		} else {
			control = 0
		}

		if control > 1 {
			return false
		}
	}

	split := strings.Split(text, "-")

	// Makes Sure That The String Is Valid In The Least e.g 1-hello
	// Returns False On 1
	if len(split) < 2 {
		return false
	}

	// Makes Sure That The String Begins With A Number And Ends With A Text e.g 1-hello-2-world
	// Returns False On 1-hello-2 Or 1-hello-2- Or hello-1-world-2-yes
	if !allTextReg.MatchString(split[len(split)-1]) || !allNumReg.MatchString(split[0]) {
		return false
	}

	//Makes Sure That String Consists Of Only Number And Texts, No AlphaNumerics
	// Returns False On 1-hell0-2-4ab
	for index, elem := range split {
		if index%2 == 0 {
			if !allNumReg.MatchString(elem) {
				return false
			}
			continue
		}
		if !allTextReg.MatchString(elem) {
			return false
		}
	}
	return true
}

/* averageumber :
Calculates The Mean Of The Numbers Contained In The Argument Passed To The Function

Parameters: string

Returns : float64

Difficulty: Easy
Used Time Of Completion: 10 Minutes
*/

func averageNumber(text string) float64 {
	text = strings.TrimSpace(text)
	split := strings.Split(text, "-")

	var (
		total    int
		numCount int
	)

	for _, elem := range split {
		if allNumReg.MatchString(elem) {
			num, _ := strconv.Atoi(elem)
			total += num
			numCount++
		}
	}

	average := float64(total) / float64(numCount)

	return float64(int(average*100)) / 100
}

/* wholeStory :
Extracts The Text In The Given Argument And Returns A Space Separated Concatenation

Parameters: string

Returns : string

Difficulty: Easy
Used Time Of Completion: 8 Minutes
*/

func wholeStory(text string) string {
	text = strings.TrimSpace(text)
	var result []string

	split := strings.Split(text, "-")

	for _, elem := range split {
		if allTextReg.MatchString(elem) {
			result = append(result, elem)
		}
	}

	return strings.Join(result, " ")
}

/* storyStats :
Returns A Summary Of The Argument Passed To The Function In The Form Of:
	Shortest Word
	Longest Word
	Average WordLength
	List Of Words In Given Argument With Same Number Of Words As Average WordLength

Parameters: string

Returns:
	string
	string
	float64
	[]string

Difficulty: Easy
Used Time Of Completion: 20 Minutes
*/

func storyStats(text string) (string, string, float64, []string) {
	text = strings.TrimSpace(text)
	split := strings.Split(text, "-")
	var (
		words           []string
		longestWord     string
		shortestWord    string
		resultList      []string
		totalWordLength int
	)

	for _, elem := range split {
		if allTextReg.MatchString(elem) {
			words = append(words, elem)
		}
	}

	shortestWord = words[0]

	for _, elem := range words {
		length := len(elem)

		if length < len(shortestWord) {
			shortestWord = elem
		} else if length > len(longestWord) {
			longestWord = elem
		}

		totalWordLength += length
	}

	average := float64(totalWordLength) / float64(len(words))

	roundedAverage := int(math.Round(average))

	for _, elem := range words {
		if len(elem) == roundedAverage {
			resultList = append(resultList, elem)
		}
	}

	return shortestWord, longestWord, float64(int(average*100)) / 100, resultList

}

/* generate:
Generates A Random Correct / Incorrect String Based On Boolean Value Passed To The Function
Parameters: bool

Returns: string

Difficulty: Medium
Estimated Time Of Completion: 20 Minutes
Used Time Of Completion: 50 Minutes
*/

func generate(valid bool) string {
	var result strings.Builder
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(12-2) + 2
	if valid {
		for i := 0; i <= length; i++ {
			if i%2 == 0 {
				num := rand.Intn(10000-1) + 1
				result.WriteString(strconv.Itoa(num))
				continue
			}
			result.WriteString("-" + randString(rand.Intn(5-1)+1, true) + "-")
		}

		// Trim Leading And Trailing "-" Characters
		returnVal := strings.TrimLeft(strings.TrimRight(result.String(), "-"), "-")

		// If Resulting String Ends With A Number, Append A String To It Before Returning
		if !testValidity(returnVal) {
			returnVal = fmt.Sprintf("%s-%s", returnVal, randString(rand.Intn(5-1)+1, true))
		}
		return returnVal

	}

	for i := 0; i <= length; i++ {
		result.WriteString(randString(rand.Intn(5-1)+1, false))

		// If By Any Slim Chance A Valid String Is Written
		// Append Another String To Render The Result Invalid
		if testValidity(result.String()) {
			result.WriteString(randString(rand.Intn(5-1)+1, false))
		}
	}

	return result.String()
}

// Helper Function For Generation Of Random Strings
func randString(length int, valid bool) string {
	var letterBytes string
	if valid {
		letterBytes = "abcdefghijklmnopqrstuvwxyz" //ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else {
		letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890-"
	}

	b := make([]byte, length)

	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}

// Exported Functions For Use By main.go

func TestValidity(text string) bool {
	return testValidity(text)
}

func AverageNumber(text string) float64 {
	return averageNumber(text)
}

func WholeStory(text string) string {
	return wholeStory(text)
}

func StoryStats(text string) (string, string, float64, []string) {
	return storyStats(text)
}

func Generate(valid bool) string {
	return generate(valid)
}

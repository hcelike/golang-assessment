package main

import (
	"assessment/assessment"
	"fmt"
	"strings"
)

func main() {
	text := "1466-boi-3784-guw-7113-rqgw-926-bht-8681-e-7143-i"
	//text := "lkhgmr2eg6-xhacqdwh1cf-mkilq9op0kj7193h9qx2cj"

	fmt.Println("Text Is Valid: ", assessment.TestValidity(text))
	shortest, longest, average, result := assessment.StoryStats(text)
	fmt.Printf("Shortest Word: %s\nLongest Word: %s\nAverage WordLength: %f\nResult: %s\n", shortest, longest, average, strings.Join(result, ","))
	fmt.Println("Whole Story: ", assessment.WholeStory(text))
	fmt.Println("Average Number: ", assessment.AverageNumber(text))
	fmt.Println("Valid Generated Text: ", assessment.Generate(true))
	fmt.Println("Invalid Generated Text: ", assessment.Generate(false))

}

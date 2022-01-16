package assessment

import (
	"testing"
)

func TestAssessment_testValidity(t *testing.T) {
	randomString := "1-golang-34-backend"

	if !testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", true, false)
	}

	randomString = "1- test-2-yes"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = " 1-please-2-hire-3-me "

	if !testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", true, false)
	}

	randomString = "1-1-ab-2-bc"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = "1--ab-2-bc-3-c"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = "22-hello-24-gophers-25"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = "hello-2-my-3-fellow-4-gophers"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = "-hi-3-there"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = "1-hello-2-testing-"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = "lkhgmr2eg6-xhacqdwh1cf-mkilq9op0kj7193h9qx2cj"

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}

	randomString = generate(true)

	if !testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", true, false)
	}

	randomString = generate(false)

	if testValidity(randomString) {
		t.Errorf("Expected: %v, Got: %v\n", false, true)
	}
	// t.Fatal("not implemented")
}

func TestAssessment_averageNumber(t *testing.T) {
	randomString := "20-hello-50-world-99-yep"

	if x := averageNumber(randomString); x != 56.33 {
		t.Errorf("Expected: %f, Got: %f\n", 56.33, x)
	}
}

func TestAssessment_wholeStory(t *testing.T) {
	text := "1-hello-2-world-3-im-4-the-5-best-6-golang-7-developer"
	expected := "hello world im the best golang developer"

	if result := wholeStory(text); result != expected {
		t.Errorf("Expected: %s, Got: %s\n", expected, result)
	}
}

func TestAssessment_storyStats(t *testing.T) {
	text := "20-hello-50-humans-99-yeah"
	expected_longest := "humans"
	expected_shortest := "yeah"
	expected_average := 5.00
	expected_result := []string{"hello"}

	shortest, longest, average, result := storyStats(text)

	if shortest != expected_shortest {
		t.Errorf("Expected: %s, Got: %s\n", expected_shortest, shortest)
	}

	if longest != expected_longest {
		t.Errorf("Expected: %s, Got: %s\n", expected_longest, longest)
	}

	if average != expected_average {
		t.Errorf("Expected: %f, Got: %f\n", expected_average, average)
	}

	if len(result) != len(expected_result) {
		t.Errorf("Expected: %d, Got: %d\n", len(expected_result), len(result))
	}

	for index, elem := range result {
		if expected_result[index] != elem {
			t.Errorf("Expected: %s At Index: %d, Got: %s\n", expected_result[index], index, elem)
		}
	}
}

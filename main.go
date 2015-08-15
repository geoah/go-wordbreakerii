package main

import "fmt"

// wordBreak receives a string and creates a map of where words start
func wordBreak(str string, dict []string) (results []string) {
	// For now force the dict to be a map to allow us to search easily
	dictMap := map[string]bool{"cat": true, "cats": true, "and": true, "sand": true, "dog": true}

	// This seems to be needed quite a lot
	strLength := len(str)
	// Create an array to map where we found words in the string
	solution := make([][]string, strLength)

	// Move from left to right on the string as i
	for i := 0; i <= strLength; i++ {
		// and then again on the leftover substring after i
		for j := i + 1; j <= strLength; j++ {
			// check if our substring is a word
			possibleWord := str[i:j]
			// and map it
			if _, ok := dictMap[possibleWord]; ok == true {
				solution[i] = append(solution[i], possibleWord)
			}
		}
	}

	fmt.Println(solution)

	return results
}

func main() {
	s := "catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}
	wordBreak(s, dict)
}

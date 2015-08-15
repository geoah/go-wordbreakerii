package main

import "fmt"

// wordBreak receives a string and creates a map of where words start
func wordBreak(str string, dict []string) (results []string) {
	// Let's convert our dict into a map to allow us to search easily for words
	dictMap := map[string]bool{}
	// Keys will be our dictionary's words
	// Values will be bool:true for no apparent reason
	for _, word := range dict {
		dictMap[word] = true
	}

	// This seems to be needed quite a lot
	strLength := len(str)
	// Create an array to map where we found words in the string
	// Keys are the character positions of the string
	// Values hold the *lengths* of the words found on each position
	solution := make([][]int, strLength)

	// Move from left to right on the string as i
	for i := 0; i <= strLength; i++ {
		// and then again on the leftover substring after i
		for j := i + 1; j <= strLength; j++ {
			// check if our substring is a word
			possibleWord := str[i:j]
			// and map it
			if _, ok := dictMap[possibleWord]; ok == true {
				solution[i] = append(solution[i], j)
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

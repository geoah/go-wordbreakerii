package main

import (
	"fmt"
	"strings"
)

// wordBreak receives a string and creates a map of where words start
func wordBreak(str string, dict []string) []string {
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

	// Create an array to store each possible sentence path/branch
	// Values of the array hold the paths
	// Values of each path hold each word's starting position on the string
	sentencePaths := [][]int{[]int{0}}
	// Create an array to store all possible sentenses
	sentences := make([]string, 0)

	for {
		// We need to create an flat array with all possible sentences
		// The first path starts from the beggining of our string which is 0
		// Starting with our first path we find the length of the words start there
		// and create a new path for each, containing the current path and the
		// starting position of the next word.
		// On the end of each loop the path list is replaced with the new paths and
		// start all over again until each path is complete.
		nextSentencePaths := [][]int{}
		for _, sentencePath := range sentencePaths {
			// Seems useful quite a bit
			sentencePathLength := len(sentencePath)
			// When going through the last word of each sentence, we basically try to
			// add the start of the next - non existing word. Which is the end of our
			// string.
			// So once the last position of a path is full length of our string, our
			// sentence is now complete and can add it to the sentenses array.
			if sentencePath[sentencePathLength-1] == strLength {
				// The sentense for this path is now complete
				// The last position in the path is fake as there is no word starting
				// there, but rather is the end of the string.
				lastPosition := sentencePathLength - 1
				// Go through all positions on the path (except the last) and construct
				// the sentence.
				temp := []string{}
				for i := 0; i < lastPosition; i++ {
					temp = append(temp, str[sentencePath[i]:sentencePath[i+1]])
				}
				// And add it to our sentences array
				sentences = append(sentences, strings.Join(temp, " "))
			} else {
				// We take the last position of the current path and find the length of
				// the words that start on this position.
				for _, j := range solution[sentencePath[sentencePathLength-1]] {
					// We create a new path containing the current path as well as the
					// position of the next word.
					// If this is the last word of the sentence, the appended position
					// will be the end of the string.
					newPath := append(sentencePath, j)
					// Add the path to the list of paths
					nextSentencePaths = append(nextSentencePaths, newPath)
				}
			}
		}
		// Once all paths have been processed, the list of new paths will be empty
		// which means that all paths have been concluded and that their sentences
		// have been added to the list
		if len(nextSentencePaths) == 0 {
			break
		} else {
			// or that we need to go through the remaining paths yet once again.
			sentencePaths = nextSentencePaths
		}
	}

	// Once all paths have been processed and no new paths were created we can
	// now return the sentences.
	return sentences
}

func main() {
	s := "catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}

	sentences := wordBreak(s, dict)

	for _, sentence := range sentences {
		fmt.Println(sentence)
	}
}

package challenge_1

import "strings"

func countDistinctWords(messages []string) int {
	words := make(map[string]bool)
	for _, sent := range messages {
		sentence := strings.Fields(sent)
		for _, word := range sentence {
			lowerWord := strings.ToLower(word)
			words[lowerWord] = true
		}
	}

	return len(words)
}

/*


Complete the countDistinctWords function using a map. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)
Copy icon
count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.

Hint
Go's strings package can be very helpful here. Specifically the Fields method and the ToLower method.



*/

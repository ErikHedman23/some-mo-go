package nested

func getNameCounts(names []string) map[rune]map[string]int {
	mapo := make(map[rune]map[string]int)

	for _, name := range names {
		n := []rune(name)
		if _, ok := mapo[n[0]]; !ok {
			mapo[n[0]] = make(map[string]int)
		}
		mapo[n[0]][name]++
	}

	return mapo
}

/*

Maps can contain maps, creating a nested structure. For example:

map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int
Copy icon
Assignment
Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names and returns a nested map. The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe
Copy icon
Creates the following nested map:

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
Copy icon
Tips
The test suite is not printing the map you're returning directly, but instead checking a few specific keys.
You can convert a string into a slice of runes as follows:
runes := []rune(word)

*/

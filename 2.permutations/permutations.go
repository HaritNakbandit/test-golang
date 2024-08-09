package main

import "fmt"

func Permutations(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	permutations := []string{}

	for i, s := range str {
		remainingStr := str[:i] + str[i+1:]
		subPermutations := Permutations(remainingStr)
		for _, p := range subPermutations {
			permutations = append(permutations, string(s)+p)
		}
	}
	return permutations
}

func main() {
	fmt.Println("Permutations [a] :", Permutations("a"))
	fmt.Println("Permutations [ab] :", Permutations("ab"))
	fmt.Println("Permutations [abc] :", Permutations("abc"))
	fmt.Println("Permutations [abcd] :", Permutations("abcd"))
}

package main

import "fmt"

func CountTheSmileyFaces(arr []string) int {
	smileys := []string{}
	for _, s := range arr {
		if len(s) == 2 && (s[0] == ':' || s[0] == ';') && (s[len(s)-1] == ')' || s[len(s)-1] == 'D') {
			smileys = append(smileys, s)
		} else if len(s) > 2 && (s[0] == ':' || s[0] == ';') && (s[1] == '-' || s[1] == '~') && (s[len(s)-1] == ')' || s[len(s)-1] == 'D') {
			smileys = append(smileys, s)
		}
	}
	return len(smileys)

}

func main() {
	fmt.Println("findTheOddInt 1 : ", CountTheSmileyFaces([]string{":)", ";(", ";}", ":-D"}))
	fmt.Println("findTheOddInt 2 : ", CountTheSmileyFaces([]string{";D", ":-(", ":-)", ";~)"}))
	fmt.Println("findTheOddInt 3 : ", CountTheSmileyFaces([]string{";]", ":[", ";*", ":$", ";-D"}))
}

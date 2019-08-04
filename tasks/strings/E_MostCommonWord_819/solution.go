package E_MostCommonWord_819

import (
	"fmt"
	"regexp"
	"strings"
)

/**
	Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.
    It is guaranteed there is at least one word that isn't banned, and that the answer is unique.
	Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.

	Example:

	Input:
		paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
		banned = ["hit"]
		Output: "ball"
	Explanation:
		"hit" occurs 3 times, but it is a banned word.
		"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
	Note that words in the paragraph are not case sensitive,
		that punctuation is ignored (even if adjacent to words, such as "ball,"),
		and that "hit" isn't the answer even though it occurs more because it is banned.
*/
func mostCommonWord(paragraph string, banned []string) string {
	parag := strings.ToLower(paragraph)
	reg, _ := regexp.Compile("[^a-z ]+")     // non a-z or non space
	parag = reg.ReplaceAllString(parag, " ") // remove symbols
	words := strings.Fields(parag)

	fmt.Println(parag)

	duplicate := make(map[string]int)

	for _, item := range words {

		if !inArray(item, banned) {
			// check if the item/element exist in the duplicate map
			_, exist := duplicate[item]

			if exist {
				duplicate[item] += 1 // increase counter by 1 if already in the map
			} else {
				duplicate[item] = 1 // else start counting from 1
			}
		}

	}

	maxCount := 0
	resultWord := ""
	for key, value := range duplicate {
		if value >= maxCount {
			maxCount = value
			resultWord = key
		}
	}
	return resultWord
}

func inArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

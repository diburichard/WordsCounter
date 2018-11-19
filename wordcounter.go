package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func SplitTextFile(txtContent string) []string {
	return strings.Fields(txtContent)
}

func FormatLowerText(txtContent string) string {
	return strings.ToLower(txtContent)
}

func RemoveSpecialCharacters(txtContent string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(txtContent, " ")
}

func CountWords(words []string) map[string]int {

	wordcountresult := make(map[string]int)
	for _, word := range words {
		wordcountresult[word]++
	}
	return wordcountresult
}

func PrintOrdened(words map[string]int) {

	keys := make([]string, 0, len(words))
	for k := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, words[k])
	}

}

func main() {

	var pathFile string
	fmt.Println("Enter the address and name of the text file")
	fmt.Scan(&pathFile)
	txt, err := ioutil.ReadFile(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	ContentFormatted := FormatLowerText(RemoveSpecialCharacters(string(txt)))
	ContentWords := SplitTextFile(ContentFormatted)
	PrintOrdened(CountWords(ContentWords))
	fmt.Scanln()
}

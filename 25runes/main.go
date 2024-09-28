package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type textAnalyzer struct {
	lineCount  int
	wordCount  int
	charCount  int
	upperCount int
	lowerCount int
}

func NewTextAnalyzer() *textAnalyzer {
	return &textAnalyzer{}
}

func (ta textAnalyzer) PrintStats() {
	fmt.Printf("line count is: %d\n", ta.lineCount)
	fmt.Printf("Word Count is: %d\n", ta.wordCount)
	fmt.Printf("Character count is: %d\n", ta.charCount)
	fmt.Printf("UpperCase letters count is: %d\n", ta.upperCount)
	fmt.Printf("LowerCase letters count is: %d\n", ta.lowerCount)
}

func (a *textAnalyzer) Analyze(text string) {
	runeCount := utf8.RuneCountInString(text)
	a.lineCount = runeCount - utf8.RuneCountInString(strings.ReplaceAll(text, "\n", ""))
	a.wordCount = len(strings.Fields(text))
	a.charCount = runeCount

	for _, r := range text {

		if unicode.IsUpper(r) {
			a.upperCount++
		} else if unicode.IsLower(r) {
			a.lowerCount++
		}
	}
}

func main() {
	text := `HELLO, ッヲギ! This is the test.
	Here's is some text with uppercase lowercase
	and even some emojis: 🏖️
	And Here's the line to count number of lines.`

	a := NewTextAnalyzer()
	a.Analyze(text)
	a.PrintStats()
}

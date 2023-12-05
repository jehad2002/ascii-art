package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func String(text string) string {
	if len(text) == 0 {
		return ""
	}
	chars := []byte(text)
	s := getArt(chars[0])
	for i := 1; i < len(chars); i++ {
		s = horizontalMerge(s, getArt(chars[i]))
	}

	return strings.Replace(strings.Join(s, "\n"), "$", " ", -1)
}

func horizontalMerge(left, right []string) []string {
	var result []string

	minLength := len(left)
	if len(right) < minLength {
		minLength = len(right)
	}

	for i := 0; i < minLength; i++ {
		mergedString := left[i] + right[i]
		result = append(result, mergedString)
	}

	return result
}

func main() {
	fmt.Println("Enter words:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(strings.TrimSpace(text), "\\n")
	asciiArt := []string{}
	for _, line := range lines {
		asciiArt = append(asciiArt, String(line))
	}
	asciiArtText := strings.Join(asciiArt, "\n")
	fmt.Println(asciiArtText)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func jehad(text string) string {
	if len(text) == 0 {
		return ""
	}
	mr := []byte(text)
	var result []string
	for i := 0; i < len(mr); i++ {
		asciim := getArt(mr[i])
		if i == 0 {
			result = asciim
		} else {
			for j := 0; j < len(result); j++ {
				result[j] += asciim[j]
			}
		}
	}
	return strings.Replace(strings.Join(result, "\n"), "$", " ", -1)
}
func main() {
	fmt.Println("Enter the string:")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(strings.TrimSpace(text), "\\n")
	asciiArt := []string{}
	for _, line := range lines {
		asciiArt = append(asciiArt, jehad(line))
	}
	asciiArtText := strings.Join(asciiArt, "\n")
	fmt.Println(asciiArtText)
}

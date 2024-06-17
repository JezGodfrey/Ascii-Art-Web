package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetFileLines(p string) []string {
	file, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Use bufio to scan lines and store them in slice lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}

func AsciiArt(t string, s string) string {
	// Checking for non-printable characters (bad request)
	for _, v := range t {
		if v < 32 || v == 127 {
			return "error"
		}
	}

	var ascii string

	// Open the file containing ascii-art
	lines := GetFileLines("asciiart/" + s + ".txt")
	if lines == nil {
		return "error"
	}

	// Convert raw newlines from input to escaped newlines
	var line []rune
	for _, v := range t {
		if v == 13 {
			continue
		}

		// Handling non-printable characters
		if v == 10 {
			line = append(line, 92)
			line = append(line, 110)
		} else {
			line = append(line, v)
		}
	}

	// Split given argument by newline combination
	words := strings.Split(string(line), "\\n")

	for _, w := range words {
		// An empty element in the slice of words indicates an extra newline
		if w == "" {
			ascii = ascii + "\n"
		} else {
			// For each character in a 'word', print each of the 8 lines corresponding to ascii-art file
			for i := 0; i < 8; i++ {
				for _, v := range w {
					ascii = ascii + lines[((int(v)-32)*9)+1+i]
				}
				ascii = ascii + "\n"
			}
		}
	}

	return ascii
}

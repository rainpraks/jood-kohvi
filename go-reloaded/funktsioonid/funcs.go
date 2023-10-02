package funktsioonid

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Hex(inputFileContents string) string {
	hexPattern := regexp.MustCompile(`(\b[0-9A-Fa-f]+\b)\s*\(hex\)`)

	updatedContents := hexPattern.ReplaceAllStringFunc(string(inputFileContents), func(match string) string {
		hexNumber := match[:len(match)-6]
		decimalValue, err := strconv.ParseInt(hexNumber, 16, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(decimalValue, 10)
	})
	return updatedContents

}

func Bin(inputFileContents string) string {
	binPattern := regexp.MustCompile(`\b([01]+)\s*\(bin\)`)
	updatedContents := binPattern.ReplaceAllStringFunc(string(inputFileContents), func(match string) string {
		binNumber := match[:len(match)-6]
		decimalValue, err := strconv.ParseInt(binNumber, 2, 64)
		if err != nil {
			return match // Return the original match if conversion fails.
		}
		return strconv.FormatInt(decimalValue, 10)
	})
	return updatedContents
}

func Up(inputFileContents string) string {

	upPattern := regexp.MustCompile(`\b\w+\b\s*\(up\)`)

	updatedContents := upPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {

		word := strings.TrimSuffix(match, " (up)")

		return strings.ToUpper(word)
	})

	return updatedContents

}

func Low(inputFileContents string) string {

	lowPattern := regexp.MustCompile(`\b\w+\b\s*\(low\)`)

	updatedContents := lowPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {

		word := strings.TrimSuffix(match, " (low)")

		return strings.ToLower(word)
	})

	return updatedContents

}

func Cap(inputFileContents string) string {

	capPattern := regexp.MustCompile(`\b\w+\b\s*\(cap\)`)

	updatedContents := capPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {

		word := strings.TrimSuffix(match, " (cap)")

		return strings.Title(word)
	})

	return updatedContents

}

func UpNum(inputFileContents string) string {
	upNumPattern := regexp.MustCompile(`\b(\w+(?: \w+)*) \(up, (\d+)\)`)

	updatedContents := upNumPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {
		matches := upNumPattern.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match
		}
		n, err := strconv.Atoi(matches[2])
		if err != nil {
			return match
		}
		words := strings.Fields(matches[1])
		start := len(words) - n //salvestab slice-i
		if start < 0 {
			start = 0
		}
		for i := start; i < len(words); i++ {
			words[i] = strings.ToUpper(words[i])
		}
		return strings.Join(words, " ")
	})

	return updatedContents
}

func LowNum(inputFileContents string) string {
	LowNumPattern := regexp.MustCompile(`\b(\w+(?: \w+)*) \(low, (\d+)\)`)

	updatedContents := LowNumPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {
		matches := LowNumPattern.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match
		}
		n, err := strconv.Atoi(matches[2])
		if err != nil {
			return match
		}
		words := strings.Fields(matches[1])
		start := len(words) - n
		if start < 0 {
			start = 0
		}
		for i := start; i < len(words); i++ {
			words[i] = strings.ToLower(words[i])
		}
		return strings.Join(words, " ")
	})

	return updatedContents
}

func CapNum(inputFileContents string) string {
	capNumPattern := regexp.MustCompile(`\b(\w+(?: \w+)*) \(cap, (\d+)\)`)

	updatedContents := capNumPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {
		matches := capNumPattern.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match
		}
		n, err := strconv.Atoi(matches[2])
		if err != nil {
			return match
		}
		words := strings.Fields(matches[1])
		start := len(words) - n
		if start < 0 {
			start = 0
		}
		for i := start; i < len(words); i++ {
			words[i] = strings.Title(words[i])
		}
		return strings.Join(words, " ")
	})

	return updatedContents
}

func FormatPunctuation(inputFileContents string) string {

	punctuationPattern := regexp.MustCompile(`(\s+)([,!.:;]+)`)

	updatedText := punctuationPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {

		trimmedMatch := strings.TrimSpace(match)

		if trimmedMatch != "" {
			fmt.Println(trimmedMatch)
			if trimmedMatch == "," {
				return trimmedMatch + " "
			}
			return trimmedMatch
		}
		return match
	})
	return updatedText
}

func FormatApostrophes(inputFileContents string) string {

	apostrophePattern := regexp.MustCompile(`'([^']+)'`)

	updatedText := apostrophePattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {

		submatches := apostrophePattern.FindStringSubmatch(match)
		if len(submatches) == 2 {

			word := strings.TrimSpace(submatches[1])
			return "'" + word + "'"
		}
		return match
	})

	return updatedText
}
func ReplaceAWithAn(inputFileContents string) string {

	aPattern := regexp.MustCompile(`\b(?i:a)\s+(?i:[aeiouh])\w*`)

	updatedText := aPattern.ReplaceAllStringFunc(inputFileContents, func(match string) string {

		word := strings.TrimSpace(match[1:])

		if len(word) > 0 && (strings.HasPrefix(strings.ToLower(word), "a") ||
			strings.HasPrefix(strings.ToLower(word), "e") ||
			strings.HasPrefix(strings.ToLower(word), "i") ||
			strings.HasPrefix(strings.ToLower(word), "o") ||
			strings.HasPrefix(strings.ToLower(word), "u") ||
			strings.HasPrefix(strings.ToLower(word), "h")) {
			return "an " + word
		}

		return match
	})

	return updatedText
}

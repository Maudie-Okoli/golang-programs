package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ProcessText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		// --- Hexadecimal conversion ---
		if words[i] == "(hex)" && i > 0 {
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = removeWord(words, i)
			i--
		}

		// --- Binary conversion ---
		if words[i] == "(bin)" && i > 0 {
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = removeWord(words, i)
			i--
		}

		// --- Single-word case ---
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = removeWord(words, i)
			i--
		}
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = removeWord(words, i)
			i--
		}
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = Capitalize(words[i-1])
			words = removeWord(words, i)
			i--
		}

		// --- Multi-word case ---
		if i+1 < len(words) {
			switch words[i] {
			case "(up,":
				n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
				words = applyMultiCase(words, i, n, "up")
				words = removeWord(words, i)
				words = removeWord(words, i)
				i--
			case "(low,":
				n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
				words = applyMultiCase(words, i, n, "low")
				words = removeWord(words, i)
				words = removeWord(words, i)
				i--
			case "(cap,":
				n, _ := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
				words = applyMultiCase(words, i, n, "cap")
				words = removeWord(words, i)
				words = removeWord(words, i)
				i--
			}
		}
	}

	result := strings.Join(words, " ")
	result = FixPunctuation(result)
	result = FixArticles(result)

	return result
}

// --- Helper functions ---

func removeWord(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// Capitalize first letter
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// Multi-word case application
func applyMultiCase(words []string, i int, n int, action string) []string {
	for j := 1; j <= n && i-j >= 0; j++ {
		switch action {
		case "up":
			words[i-j] = strings.ToUpper(words[i-j])
		case "low":
			words[i-j] = strings.ToLower(words[i-j])
		case "cap":
			words[i-j] = Capitalize(words[i-j])
		}
	}
	return words
}

// Fix punctuation spacing
func FixPunctuation(text string) string {
	re := regexp.MustCompile(`\s+([.,!?:;])`)
	text = re.ReplaceAllString(text, "$1")

	re2 := regexp.MustCompile(`([.,!?:;])([^\s.,!?:;])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	// Special cases
	text = strings.ReplaceAll(text, ". . .", "...")
	text = strings.ReplaceAll(text, "! ?", "!?")

	// Fix quotes
	re3 := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = re3.ReplaceAllString(text, "'$1'")

	return text
}

// Fix a -> an before vowels or h
func FixArticles(text string) string {
	re := regexp.MustCompile(`\b([Aa]) (\w+)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, " ")
		word := parts[1]
		if len(word) > 0 && isVowelOrH(rune(word[0])) {
			return "an " + word
		}
		return match
	})
}

func isVowelOrH(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h'
}

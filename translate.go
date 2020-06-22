package main

import "strings"

var vowels = map[rune]rune{
	'a': '1',
	'e': '2',
	'i': '3',
	'o': '4',
	'u': '5',
}

var numbers = map[rune]rune{
	'1': 'a',
	'2': 'e',
	'3': 'i',
	'4': 'o',
	'5': 'u',
}

func Encode(in string) string {
	var encodedWords []string
	words := strings.Split(in, " ")

	for _, word := range words {
		word = strings.Map(func(r rune) rune {
			if u, ok := vowels[r]; ok {
				return u
			}
			return r
		}, word)

		encodedWords = append(encodedWords, word)
	}

	return strings.Join(encodedWords, " ")
}

func Decode(in string) string {
	var decodedWords []string
	words := strings.Split(in, " ")

	for _, word := range words {
		word = strings.Map(func(number rune) rune {
			if vowel, ok := numbers[number]; ok {
				return vowel
			}
			return number
		}, word)

		decodedWords = append(decodedWords, word)
	}

	return strings.Join(decodedWords, " ")
}

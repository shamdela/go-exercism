// Package foodchain which generates the lyrics of the song 'I Know an Old Lady Who Swallowed a Fly'.
package foodchain

import (
	"fmt"
	"strings"
)

// Constant declarations
const testVersion = 3

var lyrics = map[int][]string{
	1: []string{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", "I don't know why she swallowed the fly. Perhaps she'll die."},
	2: []string{"spider", "It wriggled and jiggled and tickled inside her.", "She swallowed the spider to catch the fly."},
	3: []string{"bird", "How absurd to swallow a bird!", "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."},
	4: []string{"cat", "Imagine that, to swallow a cat!", "She swallowed the cat to catch the bird."},
	5: []string{"dog", "What a hog, to swallow a dog!", "She swallowed the dog to catch the cat."},
	6: []string{"goat", "Just opened her throat and swallowed a goat!", "She swallowed the goat to catch the dog."},
	7: []string{"cow", "I don't know how she swallowed a cow!", "She swallowed the cow to catch the goat."},
	8: []string{"horse", "She's dead, of course!", ""},
}

// Verse function to generate a single verse based on an int that denotes the verse's order of appearance.
func Verse(i int) string {
	sentence := fmt.Sprintf("I know an old lady who swallowed a %s.\n", lyrics[i][0])
	sentence = sentence + lyrics[i][1]
	for k := i; k > 0; k-- {
		if i != 1 && i != 8 {
			sentence = sentence + "\n" + lyrics[k][2]
		}
	}
	return sentence

}

// Verses function to generate two verses based on the orders given.
func Verses(a, b int) string {
	verses := []string{}
	for i := a; i <= b; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

// Song function to generate entire song.
func Song() string {
	song := Verse(1)
	for i := 2; i <= 8; i++ {
		song = song + "\n\n" + Verse(i)
	}
	return song
}

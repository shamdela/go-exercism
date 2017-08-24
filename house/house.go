// Package house to output the lyrics to 'The house that Jack built'
package house

import (
	"bytes"
	"fmt"
)

// Constant declarations
const (
	testVersion        = 1
	songStart   string = "This is"
)

var songLyrics = []string{
	"the house that Jack built.",
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to"}

// Song function to output entire Twelve days song
func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= len(songLyrics); i++ {
		// Write each verse to a Buffer
		buffer.WriteString(Verse(i))

		// Add on 2 newlines, to split all verses except at end of loop.
		if i < len(songLyrics) {
			buffer.WriteString("\n\n")
		}
	}
	return buffer.String()
}

// Verse function to output a verse of the song, given the day. Day number passed in as param.
func Verse(verse int) string {
	return fmt.Sprint(songStart, " ", getVerse(verse))
}

// getVerse function recursively calls itself to formulate the lines of a verse, given the day. Day number passed in as param.
func getVerse(verse int) string {
	// Recursively calls itself until verses count down to 0, appending each verse as it goes
	for verse > 1 {
		verse--
		return fmt.Sprint(songLyrics[verse], " ", getVerse(verse))
	}
	return fmt.Sprint(songLyrics[0])
}

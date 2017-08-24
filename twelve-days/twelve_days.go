// Package to output the lyrics to 'The Twelve Days of Christmas'
package twelve

import (
	"fmt"
	"strings"
)

// Constant declarations.
const (
	testVersion = 1
	fullStop    = "."
	songLine    = "On the %s day of Christmas my true love gave to me"
)

var gifts = []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds",
	"five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking",
	"nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth",
	"eleventh", "twelfth"}

// Song function to output entire Twelve days song
func Song() string {
	var lineStart, songString string
	lineEndings := make([]string, 1, 12) // slice of string to store each day's line ending

	for i := 0; i <= len(days)-1; i++ {
		// Swap each day into song string
		lineStart = fmt.Sprintf(songLine+",", days[i])

		// Pre-pending to slice. Add each line-ending to start of slice. eg - LIFO
		lineEndings = append([]string{fmt.Sprintf(" %s", gifts[i])}, lineEndings...)

		// Using Join to separate each string in the slice (except last one), with a comma
		// Using Trim to remove square brackets from the printed slice
		songString += fmt.Sprintf("%s%v.\n", lineStart,
			strings.Trim(strings.Join(lineEndings[:len(lineEndings)-1], ","), "[]"))

		// Pre-pending 'and' to first element, in first iteration of loop, ONLY
		if i == 0 {
			lineEndings[0] = " and" + lineEndings[0]
		}
	}
	return songString
}

// Verse function to output a line of song, given the day. Day number passed in as param.
func Verse(day int) string {
	var lineEnd string

	// Swap day string into song string
	lineStart := fmt.Sprintf(songLine, days[day-1])
	for i := day - 1; i >= 0; i-- {
		if day > 1 && i == 0 {
			// Handles pre-pendng the 'and' for the first day - 'a Partridge in a Pear Tree'
			lineEnd += fmt.Sprintf(", and %s", gifts[i])
		} else {
			// if day is not the first day, format string as normal - no pre-pending of 'and'
			lineEnd += fmt.Sprintf(", %s", gifts[i])
		}
	}
	return lineStart + lineEnd + fullStop
}

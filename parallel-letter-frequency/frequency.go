// Package letter which counts the frequency of letters in texts using parallel computation
package letter

import "sync"

// Constant declarations
const testVersion = 1

// FreqMap type which represents a map of the total frequency of each letter in a list of texts
type FreqMap map[rune]int

var wg sync.WaitGroup

// Frequency function to sequentially count letter frequencies in a single text.
func Frequency(s string) FreqMap {

	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency function to count the frequency of letters in texts using parallel computation
func ConcurrentFrequency(input []string) FreqMap {

	freqMapChannel := make(chan FreqMap, len(input)) // Channel to store results of each goroutine
	combined := FreqMap{}
	wg.Add(len(input)) // Adds a WaitGroup of length of input string

	// Call goroutines to handle each element of the input string slice. Return results to freqMapChannel.
	for i := 0; i < len(input); i++ {
		go func(text string) {
			freqMapChannel <- Frequency(text)
			wg.Done()
		}(input[i])
	}

	wg.Wait() // Wait until channels are finished processing
	close(freqMapChannel)

	// Combine results of each goroutine into a FreqMap object and return this value.
	for element := range freqMapChannel {
		for char, count := range element {
			combined[char] += count
		}
	}

	return combined
}

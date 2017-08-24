// Package gigasecond to calculate the moment when someone has lived for 10^9 seconds.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// Function to add a gigasecond to a given time
func AddGigasecond(aTime time.Time) time.Time {
	// Create a time.Duration of 1,000,000,000 seconds (A gigasecond).
	var durationSeconds time.Duration = 1000000000 * time.Second

	// Add this gigasecond to time passed in, in method arg and return
	return aTime.Add(durationSeconds)
}



// A Package that implements a clock and handles times without dates.
package clock

import (
	"fmt"
	"math"
)

// The value of testVersion here must match `targetTestVersion` in the file clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	var extraHours, newHours int

	extraHours = c.minute / 60
	newHours = (extraHours + c.hour) % 24
	c.hour = newHours
	if newHours < 0 {
		c.hour += 24
	}
	if c.minute < 0 {
		c.hour -= int(math.Abs(float64(newHours)))
	}

	if c.minute % 60 >= 0 {
		c.minute %= 60
	} else {
		c.minute %= 60
		c.minute += 60
	}


	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	var extraHours int

	extraHours = minutes / 60
	fmt.Println(extraHours)
	return c
}

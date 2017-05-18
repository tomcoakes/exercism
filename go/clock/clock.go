package clock

import (
	"fmt"
	"math"
)

const testVersion = 4
const minutesInADay = 1440

type Clock int

func New(hour, minute int) Clock {
	hoursInMinutes := hour * 60
	timeInMinutes := keepClocksWithin24Hours(hoursInMinutes + minute)

	return Clock(timeInMinutes)
}

func (c Clock) String() string {
	hours := int(math.Floor(float64(int(c) / 60))) % 24
	minutes := int(c) % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (c Clock) Add(minutes int) Clock {
	currentTimeInMinutes := int(c)
	newTimeInMinutes := keepClocksWithin24Hours(currentTimeInMinutes + minutes)

	return Clock(newTimeInMinutes)
}

func keepClocksWithin24Hours(timeInMinutes int) int {
	for timeInMinutes < 0 {
		timeInMinutes = timeInMinutes + minutesInADay
	}
	for timeInMinutes > minutesInADay {
		timeInMinutes = timeInMinutes - minutesInADay
	}

	return timeInMinutes
}
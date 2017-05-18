package gigasecond

import (
	"time"
)

const testVersion = 4
const gigaSecond = 1000000000

func AddGigasecond(t time.Time) time.Time {
	previousDateTimeInSeconds := t.Unix()
	newDateTimeInSeconds := previousDateTimeInSeconds + gigaSecond
	newDateTimeStamp := time.Unix(newDateTimeInSeconds, 0)
	return newDateTimeStamp
}

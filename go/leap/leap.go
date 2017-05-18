package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	if isDivisibleBy(4, year) && (!isDivisibleBy(100, year) || isDivisibleBy(400, year)) {
		return true
	}
	return false
}

func isDivisibleBy(divisor int, year int) bool {
	return year%divisor == 0
}

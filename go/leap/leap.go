package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	if isDivisibleByFour(year) && (!isDivisibleByOneHundred(year) || isDivisibleByFourHundred(year)) {
		return true
	}
	return false
}

func isDivisibleByFour(year int) bool {
	return year%4 == 0
}

func isDivisibleByOneHundred(year int) bool {
	return year%100 == 0
}

func isDivisibleByFourHundred(year int) bool {
	return year%400 == 0
}

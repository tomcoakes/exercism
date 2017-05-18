package accumulate

const testVersion = 1

func Accumulate(collection []string, function func(string) string) []string {
	var outputCollection = []string{}

	for _, value := range collection {
		outputCollection = append(outputCollection, function(value))
	}
	return outputCollection
}

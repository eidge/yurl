package request

func isStringInArray(expectedValue string, array []string) bool {
	for _, value := range array {
		if value == expectedValue {
			return true
		}
	}
	return false
}

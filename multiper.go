package multiper

func Multiplier(values ...int) (int, string) {
	result := 0
	for i := range values {
		for j := i + 1; j < len(values); j++ {
			result = result + values[i]*values[j]
		}
	}
	string := "lco"
	return result, string
}

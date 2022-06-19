package util

func Contains(seq []string, value string) bool {
	for _, v := range seq {
		if v == value {
			return true
		}
	}
	return false
}

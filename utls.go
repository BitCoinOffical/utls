package utls

func InStrSlice(a []string, x string) bool {
	for _, i := range a {
		if x == i {
			return true
		}
	}
	return false
}
func InIntSlice(a []int, x int) bool {
	for _, i := range a {
		if x == i {
			return true
		}
	}
	return false
}

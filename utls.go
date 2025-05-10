package utls

func Contains(a []string, x string) bool {
	for _, i := range a {
		if x == i {
			return true
		}
	}
	return false
}

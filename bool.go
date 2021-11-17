package _mock

func Bool() bool {
	if Int64Range(0, 1) == 0 {
		return false
	}
	return true
}

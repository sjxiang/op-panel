package helper


func If(ok bool, a, b interface{}) interface{} {
	if ok {
		return a
	}

	return b
}
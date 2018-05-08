package redisPagination

func trinocular(condition bool, first, second interface{}) interface{} {
	if condition {
		return first
	} else {
		return second
	}
}

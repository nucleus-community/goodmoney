package good_money

func absInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

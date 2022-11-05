package database

import "strconv"

func stringToUint(s string) (uint, error) {
	ans, err := strconv.ParseUint(s, 10, 32)
	return uint(ans), err
}

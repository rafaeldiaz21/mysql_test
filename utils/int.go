package utils

import (
	"math/rand"
	"time"
)

func RandomNum(num int64) int {
	rand.Seed(time.Now().UTC().UnixNano() + num)
	return rand.Intn(6)
}

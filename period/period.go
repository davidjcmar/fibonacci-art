package period

import (
	"math/big"
)

const (
	sliceCap = 100 // Initial and additional capacity for slice
)

/* fib Returns closure for generating fibonacci numbers */
func fib() func() *big.Int {
	a, b := big.NewInt(int64(0)), big.NewInt(int64(1))
	return func() *big.Int {
		ret := a
		a, b = b, big.NewInt(int64(0)).Add(a, b)
		return ret
	}
}

/* growSlice Expands capacity for slice by sliceCap const */
func growSlice(s []uint64) []uint64 {
	newCap := cap(s) + sliceCap
	s1 := make([]uint64, cap(s), newCap)
	copy(s1, s)
	return s1
}

/* checkPeriod Determines if entire period has been captured */
func checkPeriod(s []uint64) bool {
	if len(s) > 3 && s[len(s)-1] == 1 && s[len(s)-2] == 1 && s[len(s)-3] == 0{
		return true
	}
	return false
}

/* Pisano Returns []uint64 containing pisano period */
func Pisano(m uint) ([]uint64, error) {
	pisanoPeriod := make([]uint64, 0, sliceCap)
	f := fib()
	for {
		if len(pisanoPeriod) == cap(pisanoPeriod) {
			pisanoPeriod = growSlice(pisanoPeriod)
		}
		nextFib := f()
		pisanoPeriod = append(pisanoPeriod, big.NewInt(int64(0)).Mod(nextFib, big.NewInt(int64(m))).Uint64())
		if checkPeriod(pisanoPeriod) {
			break
		}
	}
	return pisanoPeriod[:len(pisanoPeriod)-3], nil
}

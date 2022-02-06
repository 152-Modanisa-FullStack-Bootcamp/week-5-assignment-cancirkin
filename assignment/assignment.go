package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sumUint64 := uint64(x) + uint64(y)
	maxUint32 := 1<<32 - 1
	if sumUint64 > uint64(maxUint32) {
		return x + y, true
	}
	return x + y, false
}

func CeilNumber(f float64) float64 {
	return math.Ceil(4*f) / 4
}

func AlphabetSoup(s string) string {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func StringMask(s string, n uint) string {
	rs := []rune(s)
	m := int(n)
	for i := 1; i < len(rs); i++ {
		if i >= m {
			rs[i] = '*'
		}
	}
	return string(rs)
}

func WordSplit(arr [2]string) string {
	wordsArr := strings.Split(arr[1], ",")
	var sentence []string
	for i, _ := range wordsArr {
		if strings.Index(arr[0], wordsArr[i]) > -1 {
			sentence = append([]string{wordsArr[i]}, sentence...)
		}
	}

	if sentence == nil {
		return "not possible"
	}
	return strings.Join(sentence, ",")
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}

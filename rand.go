package structdefault

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func rangeIt(org string) (int64, error) {
	var min, max int64
	var err error
	var rangeExp = regexp.MustCompile(`(?P<min>\d+),(?P<max>\d+)`)
	match := rangeExp.FindStringSubmatch(org)
	result := make(map[string]string)
	for i, name := range rangeExp.SubexpNames() {
		result[name] = match[i]
	}
	if min, err = strconv.ParseInt(result["min"], 10, 0); err != nil {
		return 0, err
	}
	if max, err = strconv.ParseInt(result["max"], 10, 0); err != nil {
		return 0, err
	}
	return randomInt64(min, max), nil

}

func randomInt64(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}

package gosdk

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func IsExisted(e map[string]struct{}, key string) bool {
	_, ok := e[key]
	if ok {
		return true
	}
	return false
}

func MergeStringSlice(s1 []string, sList ...[]string) []string {
	for _, s := range sList {
		s1 = append(s1, s...)
	}
	return s1
}

func TrimInList(word string, sep string, trimList map[string]struct{}) string {
	splitWordList := strings.Split(word, sep)
	if _, ok := trimList[splitWordList[1]]; !ok {
		return word
	}

	return strings.TrimPrefix(word, sep+splitWordList[1])
}

func BoolAdr(b bool) *bool {
	return &b
}

func StringAdr(in string) *string {
	return &in
}

func IntAdr(in int) *int {
	return &in
}

func UUIDAdr(in uuid.UUID) *uuid.UUID {
	return &in
}

func GetCurrentTimePtr() *time.Time {
	tmp := time.Now()
	return &tmp
}

func GetCurrentYear2Digit() int {
	return (time.Now().Year() + 543) % 100
}

func CalYearFromID(sid string) (string, error) {
	if len(sid) != 10 {
		return "", errors.New("Invalid student id")
	}

	yearIn, err := strconv.Atoi(sid[:2])
	if err != nil {
		return "", errors.New("Invalid student id")
	}

	studYear := GetCurrentYear2Digit() - yearIn + 1
	if studYear <= 0 {
		return "", errors.New("Invalid student ID")
	}

	if time.Now().YearDay() < 213 {
		studYear = studYear - 1
	}

	return strconv.Itoa(studYear), nil
}

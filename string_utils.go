package go_commons

import (
	cRand "crypto/rand"
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cRand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func BoolToString(input bool) string {
	return strconv.FormatBool(input)
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func Int64ToString(input int64) string {
	return strconv.FormatInt(input, 10)
}

func StringToInt64(input string) (output int64, err error) {
	if input == "" {
		output = 0
		err = nil
	} else {
		output, err = strconv.ParseInt(input, 0, 64)
	}
	return
}

func IsStringContainsAnyKeywords(s string, keywords []string) bool {
	contain := false
	for i := range keywords {
		if strings.Contains(s, keywords[i]) {
			contain = true
			break
		}
	}
	return contain
}

func StringTrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func IsStringEmpty(s string) bool {
	return s == ""
}

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

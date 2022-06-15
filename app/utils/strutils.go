package utils

import (
	crypto "crypto/rand"
	math "math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                                                      // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits                                     // # of letter indices fitting in 63 bits
)

// RandomAlphaString from math module
func RandomAlphaString(n int) string {
	var src = math.NewSource(time.Now().UnixNano())

	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

//SecureRandomAlphaString of len n by crypt
func SecureRandomAlphaString(length int) (str string) {
	defer func() {
		if err := recover(); err != nil {
			str = RandomAlphaString(length)
		}
	}()

	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i++
		}
	}

	str = string(result)
	return
}

// secureRandomBytes returns the requested number of bytes using crypto/rand
func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	if _, err := crypto.Read(randomBytes); err != nil {
		panic(err)
	}
	return randomBytes
}

func ParseFullName(name string) (string, string, string) {

	name = strings.TrimSpace(name)
	re := regexp.MustCompile(`/\s\s+/g`)

	s := re.ReplaceAllString(name, " ")

	if name == "" {
		return "", "", ""
	}

	nameArr := strings.Split(s, " ")

	size := len(nameArr)

	if size == 1 {
		return nameArr[0], "", ""
	} else if size == 2 {
		return nameArr[0], "", nameArr[1]
	}

	firstName := strings.TrimSpace(nameArr[0])
	middleName := strings.TrimSpace(nameArr[1])
	lastName := strings.Join(nameArr[2:], " ")

	return firstName, middleName, lastName
}

func NormalizeNFD(text string) string {
	// t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, text)
	return result
}

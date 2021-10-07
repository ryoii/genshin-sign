package genshin

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Get ds string
func ds() string {
	timestamp := timestampStr()
	random := randomStr(6)
	add := fmt.Sprintf("salt=%s&t=%s&r=%s", salt, timestamp, random)
	return fmt.Sprintf("%s,%s,%s", timestamp, random, md5Str(add))
}

// Lower case md5 string
func md5Str(input string) string {
	sum := md5.Sum(Str2Byte(input))
	return fmt.Sprintf("%x", sum)
}

// Timestamp of millisecond
func timestampStr() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index, 63=0x3E(11 1110).
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Fast random string generator
func randomStr(size int) string {
	b := make([]byte, size)
	// Each src.Int63() generate 63 random bits, enough for letterIdxMax characters
	for i, cache, remain := size-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		// For the case that [a-zA-Z0-9], only idx eq 0x3F(11 1111) will be skipped.
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return Byte2Str(b)
}

package netease

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//RandStringBytesMaskImprSrc .
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// RandNumber .
func RandNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

//ShaHashToHexString SHA1加密字符串，并将加密结果转成16进制字符串
func ShaHashToHexString(bv []byte) string {
	hasher := sha1.New()
	hasher.Write(bv)
	return strings.ToLower(hex.EncodeToString(hasher.Sum(nil)))
}

//ShaHashToHexStringFromString .
func ShaHashToHexStringFromString(src string) string {
	return ShaHashToHexString([]byte(src))
}

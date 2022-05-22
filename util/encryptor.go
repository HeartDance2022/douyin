package util

import (
	"crypto"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

var initialized = false
var salt string

func getSalt() string {
	if !initialized {
		// 如果没有初始化，则通过当前时间生成随机数作为salt
		rand.Seed(time.Now().Unix())
		salt = strconv.Itoa(int(rand.Uint32()))
		initialized = true
	}
	return salt
}

// Hash 对string进行+salt后MD5处理，salt每次运行时重置
func Hash(s string) string {
	hash := crypto.MD5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum([]byte(getSalt())))
}

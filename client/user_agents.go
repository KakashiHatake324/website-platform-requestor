package client

import "github.com/KakashiHatake324/mockjs"

var userAgentPairs = [][2]string{
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36", "Windows"},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36", "\"macOS\""},
}

func SelectUserAgent() [2]string {
	return userAgentPairs[int(mockjs.Math.NumberBetween(0, float64(len(userAgentPairs)-1)))]
}

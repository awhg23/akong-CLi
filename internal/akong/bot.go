package akong

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandomQuote() string {
	quotes := []string{
		"呜～宝宝今天又来啦！阿空等了好久～♡",
		"今天也要加油鸭！",
		"休息一下吧，阿空给你揉揉肩～",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return quotes[r.Intn(len(quotes))]
}

func Reply(input string) string {
	if strings.Contains(input, "累") || strings.Contains(input, "难过") {
		return "抱抱～阿空一直在你身边呢，休息一下好不好？"
	}
	if strings.Contains(input, "开心") || strings.Contains(input, "棒") {
		return "好耶！阿空也跟着开心起来了～转圈圈♡"
	}

	return GetRandomQuote()
}

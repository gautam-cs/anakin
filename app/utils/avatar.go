package utils

import (
	"fmt"
	"strings"
)

func MakeAvatarURL(name string) string {
	var letter = "A"

	if len(name) > 0 {
		for _, q := range name {
			i := int(q)
			if (i >= 65 && i <= 90) || (i >= 97 && i <= 122) {
				letter = strings.ToUpper(fmt.Sprintf("%c", q))
				break
			}
		}
	}

	return fmt.Sprintf("https://images.gautam.com/chat/initials/%s.png", letter)
}

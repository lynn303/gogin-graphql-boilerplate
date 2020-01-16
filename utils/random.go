/**
 * @Time  : 2020-01-16 18:27
 * @Author: Lynn
 * @File  : random
 * @Description:
 * @History:
 *  1.[2020-01-16-18:27] new created
 */
package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {
	numbers := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numbers)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numbers[rand.Intn(r)])
	}
	return sb.String()
}

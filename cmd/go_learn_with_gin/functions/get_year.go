package functions

import "time"

func GetYear() int {
	return time.Now().Year()
}

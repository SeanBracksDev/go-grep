package colour

import "strconv"

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func Colour(input interface{}, colour string) []byte {
	switch v := input.(type) {
	case string:
		return append([]byte(colour), append([]byte(v), []byte(Reset)...)...)
	case int:
		return append([]byte(colour), append([]byte(strconv.Itoa(v)), []byte(Reset)...)...)
	case []byte:
		return append([]byte(colour), append(v, []byte(Reset)...)...)
	default:
		return []byte{}
	}
}

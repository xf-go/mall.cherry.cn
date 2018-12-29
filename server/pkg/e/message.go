package e

var Messages = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

func GetMessage(code int) string {
	msg, ok := Messages[code]
	if ok {
		return msg
	}
	return Messages[ERROR]
}

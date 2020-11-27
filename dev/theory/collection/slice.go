package collection

func NewSlice(oldSlice []string, elements ...string) []string {
	// append는 새로운 slice를 만든다.
	return append(oldSlice, elements...)
}

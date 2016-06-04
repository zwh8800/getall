package util

var id = 0

func GenerateId() int {
	id++
	return id
}

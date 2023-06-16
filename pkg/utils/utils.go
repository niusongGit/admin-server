package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func GetOffset(pageSize, page int) (offset int) {
	offset = -1
	if pageSize != -1 && page != -1 {
		if page == 0 {
			page = 1
		}

		if pageSize == 0 {
			pageSize = 10
		}
		offset = (page - 1) * pageSize
	}
	return offset
}

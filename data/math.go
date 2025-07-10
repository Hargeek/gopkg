package data

import (
	r "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(length int) (int, error) {
	if length > 10 {
		return 0, fmt.Errorf("length must be less than or equal to 10")
	}
	rand.Seed(time.Now().UnixNano())
	randomNumber := strconv.Itoa(rand.Intn(9) + 1) // First digit between 1 and 9
	for i := 0; i < length-1; i++ {                // Generate the remaining digits
		randomNumber += strconv.Itoa(rand.Intn(10))
	}
	result, err := strconv.Atoi(randomNumber)
	if err != nil {
		return 0, err
	}
	return result, nil

}

func GenerateRandomID(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := r.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:n], nil
}

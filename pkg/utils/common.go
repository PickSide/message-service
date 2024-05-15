package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func GenerateRandomUsername(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	base64Encoded := base64.StdEncoding.EncodeToString(randomBytes)

	return base64Encoded, nil
}

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var result string
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		result += string(charset[randomIndex.Int64()])
	}

	return result, nil
}

func ConvertStringsToUint64Array(strArray []string) ([]uint64, error) {
	var uintArray []uint64

	for _, str := range strArray {
		uintVal, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		uintArray = append(uintArray, uintVal)
	}

	return uintArray, nil
}

func ConvertUint64ToStringsArray(arr []uint64) []string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = strconv.FormatUint(v, 10)
	}

	return strArr
}

func Uint64Includes(unint64Array []uint64, value uint64) bool {
	for _, unit := range unint64Array {
		if unit == value {
			return true
		}
	}
	return false
}

func StringToUint64(s string) (uint64, error) {
	result, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func StringSliceToUint64(strings []string) ([]uint64, error) {
	uints := make([]uint64, len(strings))
	for i, s := range strings {
		val, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error converting string to uint64 at index %d: %w", i, err)
		}
		uints[i] = val
	}
	return uints, nil
}
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func CamelToSnake(s string) string {
	var re = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")

	snake = strings.ToLower(snake)

	return snake
}

func IsEmail(str string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(str)
}

package helper

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// RandomStringBase64 function for random string and base64 encoded
func RandomStringBase64(length int) string {
	rb := make([]byte, length)
	_, err := rand.Read(rb)

	if err != nil {
		return ""
	}
	rs := base64.URLEncoding.EncodeToString(rb)

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		return ""
	}

	return reg.ReplaceAllString(rs, "")
}

// RandomNumber generate random number
func RandomNumber(length int) int {

	CHARS := "0123456789"
	rand.Seed(time.Now().UTC().UnixNano())
	charsLength := len(CHARS)
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = string(CHARS[rand.Intn(charsLength)])
	}

	return cast.ToInt(strings.Join(result, ""))
}

// GetEnv read `ENV` variable from os system
func GetEnv() (env string) {
	env = os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	return env
}

// GetStringValue return string value from SQL null string type
func GetStringValue(v sql.NullString) string {
	var value string
	if v.Valid {
		value = v.String
	}
	return value
}

// GetInt64Value return int64 value from SQL null int64 type
func GetInt64Value(v sql.NullInt64) int64 {
	var value int64
	if v.Valid {
		value = v.Int64
	}
	return value
}

// GenerateRandomToken function for generating random token
func GenerateRandomToken(n int) string {
	var letterBytes = "0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// HashPassword hash password
func HashPassword(pwd string) string {

	h := md5.New()
	io.WriteString(h, pwd)
	password := fmt.Sprintf("%x", h.Sum(nil))
	return password
}

// AvatarByName avatar by name
func AvatarByName(name string) string {
	return fmt.Sprintf("https://ui-avatars.com/api/?background=f15321&color=fff&name=%s&size=512&length=1", name)
}

// DeleteFile function for deleting file by path
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// IsValidUrl tests a string to determine if it is a url or not.
func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

// IntSliceToString : convert
func IntSliceToString(numbers []int64) string {
	var numbersString []string
	for _, id := range numbers {
		numbersString = append(numbersString, strconv.FormatInt(id, 10))
	}
	return strings.Join(numbersString, ",")
}

// InSliceString :
func InSliceString(val string, source []string) bool {
	for _, v := range source {
		if v == val {
			return true
		}
	}

	return false
}

// InSliceInt64 :
func InSliceInt64(val int64, source []int64) bool {
	for _, v := range source {
		if v == val {
			return true
		}
	}

	return false
}

// UniqueInt :
func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetGuestEmail :
func GetGuestEmail(phone string) string {
	return fmt.Sprintf("%s@guest.com", phone)
}

// CheckGuestEmail :
func CheckGuestEmail(email string) bool {
	return strings.Contains(email, "guest.com")
}

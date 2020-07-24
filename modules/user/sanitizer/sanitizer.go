package sanitizer

import (
	"errors"
	"strconv"

	"github.com/labstack/echo"
)

// GetProfile - function for validating user id
func GetProfile(c echo.Context) (int, error) {

	uid, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return 0, errors.New("user id must be numeric")
	}

	if uid == 0 {
		return 0, errors.New("user id is invalid")
	}

	return uid, nil
}

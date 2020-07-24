package sanitizer

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/auth/model"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Login  validate login params
func Login(c echo.Context) (p *model.LoginRequest, err error) {

	if err := c.Bind(&p); err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "sanitizer.Login", "bind_request")
		return nil, errors.New("Invalid request parameter")
	}
	if len(p.Email) == 0 {
		return nil, errors.New("Email address is required")
	}
	p.Email = strings.ToLower(p.Email)

	if len(p.Password) == 0 {
		return nil, errors.New("Password is required")
	}

	newPass, err := base64.StdEncoding.DecodeString(p.Password)
	if err == nil {
		p.Password = string(newPass)
	}

	check := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString
	if !check(p.Email) {
		return nil, errors.New("Email address is invalid")
	}
	return p, nil
}

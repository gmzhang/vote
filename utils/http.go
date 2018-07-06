package utils

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"time"
)

//ForwardGet is helper function for GET backend url base on request header infomation
func ForwardGet(url string, c echo.Context) (body []byte, setCookies string, err error) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Referer", c.Request().Referer())
	req.Header.Set("X-Forwarded-For", c.RealIP())
	req.Header.Set("User-Agent", c.Request().UserAgent())
	req.Header.Set("Cookie", c.Request().Header.Get("Cookie"))

	resp, err := client.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	setCookies = resp.Header.Get("Set-Cookie")
	body, err = ioutil.ReadAll(resp.Body)
	return
}

//CreateCookie is a helper function for quickly get a cookie in a line code
func CreateCookie(name, value string, expire time.Time, path string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expire
	cookie.Path = path
	return cookie
}

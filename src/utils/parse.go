package utils

import (
	"io"
	"net/http"
	"regexp"
)

// Split &&&START&&& in the login sign of response text
func ParseResponse(res *http.Response) []byte {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	return body[11:]

}

// Get serviceToken
func ParseToekn(source string) string {
	c := regexp.MustCompile("serviceToken=(.*?);")
	s := c.FindStringSubmatch(source)
	return s[len(s)-1]
}

func ParseVolume(source string) string {
	c := regexp.MustCompile("\"volume\":(.*?),")
	s := c.FindStringSubmatch(source)
	return s[len(s)-1]
}

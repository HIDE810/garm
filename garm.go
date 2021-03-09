// Package garm provides a converting function of ARM codes.
package garm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

func convertCode(body string, mode string) (string, error) {

	url := "https://armconverter.com/api/convert"

	header := map[string]string{
		"Host":            "armconverter.com",
		"Content-Length":  strconv.Itoa(len(body)),
		"Content-Type":    "application/json",
		"Accept":          "*/*",
		"Accept-Encofing": "gzip, deflate, br",
		"Connection":      "keep-alive",
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", errors.Wrap(err, "could not generate http request")
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return "", errors.Wrap(err, "request failed")
	}

	defer resp.Body.Close()

	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "could not read the body")
	}

	result := gjson.Get(string(json), mode+".1").String()

	if gjson.Get(string(json), mode+".0").String() == "false" {
		return "", errors.New(result)
	}

	return result, nil
}

// ArmToHex converts Arm assembly to hexadecimal.
func ArmToHex(code string) (string, error) {
	body := `{"asm":"` + code + `","offset":"","arch":"armbe"}`
	mode := "hex.armbe"

	result, err := convertCode(body, mode)
	if err != nil {
		return "", errors.Wrap(err, "could not convert the code")
	}
	return result, nil
}

// HexToArm converts hexadecimal to Arm assembly.
func HexToArm(code string) (string, error) {
	body := `{"hex":"` + code + `","offset":"","arch":"armbe"}`
	mode := "asm.armbe"

	result, err := convertCode(body, mode)
	if err != nil {
		return "", errors.Wrap(err, "could not convert the code")
	}
	return result, nil
}

// ThumbToHex converts Thumb assembly to hexadecimal.
func ThumbToHex(code string) (string, error) {
	body := `{"asm":"` + code + `","offset":"","arch":"thumbbe"}`
	mode := "hex.armbe"

	result, err := convertCode(body, mode)
	if err != nil {
		return "", errors.Wrap(err, "could not convert the code")
	}
	return result, nil
}

// HexToThumb converts hexadecimal to Thumb assembly.
func HexToThumb(code string) (string, error) {
	body := `{"hex":"` + code + `","offset":"","arch":"thumbbe"}`
	mode := "hex.armbe"

	result, err := convertCode(body, mode)
	if err != nil {
		return "", errors.Wrap(err, "could not convert the code")
	}
	return result, nil
}

package pdc

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func newHttpClient(to int) *http.Client {
	// Set timeout to 20 if no timeout is specified
	if to == 0 {
		to = 20
	}

	return &http.Client{
		Timeout: time.Duration(to) * time.Second,
	}
}

func postRequestForm(baseUrl, path, jSessionId string, urlValues url.Values) (*http.Response, error) {
	method := "POST"

	payload := strings.NewReader(urlValues.Encode())

	req, err := http.NewRequest(method, baseUrl+path, payload)
	if err != nil {
		return nil, err
	}

	if jSessionId != "" {
		ck := &http.Cookie{
			Name:  "JSESSIONID",
			Value: jSessionId,
		}

		req.AddCookie(ck)
	}

	client := newHttpClient(0)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != http.StatusOK {
		return nil, generateError(res)
	}

	return res, nil
}

func postRequest(baseUrl, path, jSessionId string) (*http.Response, error) {
	method := "POST"

	req, err := http.NewRequest(method, baseUrl+path, nil)
	if err != nil {
		return nil, err
	}

	if jSessionId != "" {
		ck := &http.Cookie{
			Name:  "JSESSIONID",
			Value: jSessionId,
		}

		req.AddCookie(ck)
	}

	client := newHttpClient(0)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != http.StatusOK {
		return nil, generateError(res)
	}

	return res, nil
}

func postRequestJsonString(baseUrl, path, jSessionId, jsonString string) (*http.Response, error) {
	method := "POST"

	payload := []byte(jsonString)

	req, err := http.NewRequest(method, baseUrl+path, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	if jSessionId != "" {
		ck := &http.Cookie{
			Name:  "JSESSIONID",
			Value: jSessionId,
		}

		req.AddCookie(ck)
	}

	client := newHttpClient(0)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return nil, generateError(res)
	}
	return res, nil
}

func decodeBody(body io.ReadCloser) (string, error) {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

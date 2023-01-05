package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type authTransport struct {
	next   http.RoundTripper
	bearer string
	http.RoundTripper
}

func (at *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", at.bearer)
	return at.next.RoundTrip(req)
}

type loginResponse struct {
	Status string
	Data   map[string]interface{}
}

func NewAuth(baseUrl, apikey, pin string, next http.RoundTripper) (*authTransport, error) {
	j, err := json.Marshal(map[string]string{
		"apikey": apikey,
		"pin":    pin,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+"/login", "application/json", bytes.NewReader(j))
	if err != nil {
		return nil, err
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response loginResponse
	err = json.Unmarshal(r, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("api failed")
	}

	token, exists := response.Data["token"]
	if !exists {
		return nil, fmt.Errorf("api failed")
	}

	t, ok := token.(string)
	if !ok {
		return nil, fmt.Errorf("datatype is wrong")
	}

	at := &authTransport{
		next:   next,
		bearer: "Bearer " + t,
	}
	return at, nil
}

package testexpect

import "encoding/json"

func ToString(expected any) (string, error) {
	data, err := json.Marshal(expected)
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil // response body data always have \n at the tail
}

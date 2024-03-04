package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func GetSub() ([]byte, int, error) {

	req_sub, err := http.NewRequest(http.MethodGet, "http://host.docker.internal:6450/", nil)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not create request")
	}

	req_sub.Header.Set("Content-Type", "application/json")

	res_sub, err := http.DefaultClient.Do(req_sub)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error making http request")
	}

	body_sub, err := ioutil.ReadAll(res_sub.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res_sub.Body.Close()

	return body_sub, res_sub.StatusCode, nil
}

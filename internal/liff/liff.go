package liff

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	endpoint    = "https://api.line.me"
	accessToken string

	UnauthorizedErr    = errors.New("unauthorized")
	LiffAppNotFoundErr = errors.New("no LIFF application(s) on this channel")
)

type App struct {
	LiffID string `json:"liffId"`
	View   *View  `json:"view"`
}

type View struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func ListApps() ([]*App, error) {
	r, err := req("GET", "/liff/v1/apps", nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusUnauthorized {
		return nil, UnauthorizedErr
	}
	if r.StatusCode == http.StatusNotFound {
		return nil, LiffAppNotFoundErr
	}

	var result struct {
		Apps []*App `json:"apps"`
	}
	json.NewDecoder(r.Body).Decode(&result)
	return result.Apps, nil
}

// Update LIFF app settings.
func Update(liffID string, view *View) error {
	r, err := req("PUT", "/liff/v1/apps/"+liffID+"/view", view)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusUnauthorized {
		return UnauthorizedErr
	}
	if r.StatusCode == http.StatusBadRequest {
		b, _ := ioutil.ReadAll(r.Body)
		return errors.New(string(b))
	}

	return nil
}

// Adds an app to LIFF. It returns liffID if add success.
func Add(view *View) (liffID string, err error) {
	r, err := req("POST", "/liff/v1/apps", struct {
		View *View `json:"view"`
	}{
		View: view,
	})
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusUnauthorized {
		return "", UnauthorizedErr
	}
	if r.StatusCode == http.StatusBadRequest {
		b, _ := ioutil.ReadAll(r.Body)
		return "", errors.New(string(b))
	}

	var result struct {
		LiffID string `json:"liffId"`
	}
	json.NewDecoder(r.Body).Decode(&result)
	return result.LiffID, nil
}

func Delete(id string) error {
	r, err := req("DELETE", "/liff/v1/apps/"+id, nil)
	if err != nil {
		return err
	}

	if r.StatusCode == http.StatusUnauthorized {
		return UnauthorizedErr
	}
	if r.StatusCode == http.StatusNotFound {
		return LiffAppNotFoundErr
	}

	return nil
}

func SetAccessToken(token string) {
	accessToken = token
}

func req(method, path string, v interface{}) (*http.Response, error) {
	body := marshal(v)

	req, err := http.NewRequest(method, endpoint+"/"+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func marshal(v interface{}) io.Reader {
	if v == nil {
		return nil
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(v)
	return &buf
}

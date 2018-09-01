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
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("unauthorized")
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("no LIFF applications on the channel.")
	}

	var result struct {
		Apps []*App `json:"apps"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Apps, nil
}

// Adds an app to LIFF. It returns liffID if add success.
func Add(view *View) (liffID string, err error) {
	b, err := json.Marshal(struct {
		View *View `json:"view"`
	}{
		View: view,
	})
	if err != nil {
		return "", err
	}
	r, err := req("POST", "/liff/v1/apps", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return "", errors.New("unauthorized")
	}
	if resp.StatusCode == http.StatusBadRequest {
		b, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(string(b))
	}

	var result struct {
		LiffID string `json:"liffId"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.LiffID, nil
}

func Delete(id string) error {
	r, err := req("DELETE", "/liff/v1/apps/"+id, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("unauthorized")
	}
	if resp.StatusCode == http.StatusNotFound {
		return errors.New("no LIFF applications on the channel.")
	}
	return nil
}

func SetAccessToken(token string) {
	accessToken = token
}

func req(method, path string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, endpoint+"/"+path, body)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r.Header.Add("Content-Type", "application/json")
	return r, nil
}

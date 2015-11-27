package kairos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	libVersion     = "1.0"
	defaultBaseURL = "http://api.kairos.com/"
	userAgent      = "go-kairos/" + libVersion
)

type (
	// A Client manages communications with the Kairos API.
	Kairos struct {
		// HTTP client
		client *http.Client

		// Base URL for API request
		BaseUrl *url.URL

		// User agent used when communicating with the Kairos API
		UserAgent string

		// App info used to save the id and app key
		AppInfo *url.Userinfo
	}

	// Kairos options
	Options struct {
		selector      string
		symmetricFill bool
	}
)

// New returns a new Kairos API client.
func New(api, appId, appKey string) (*Kairos, error) {
	if api == "" {
		api = defaultBaseURL
	}

	if !strings.HasSuffix(api, "/") {
		api += "/"
	}

	baseUrl, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	k := &Kairos{
		client:    http.DefaultClient,
		BaseUrl:   baseUrl,
		UserAgent: userAgent,
		AppInfo:   url.UserPassword(appId, appKey),
	}

	return k, nil
}

// Enroll an image
func (k Kairos) Enroll(image string, subjectId, galleryName string, options ...Options) error {
	// TODO
	return nil
}

// Recognize an image
func (k Kairos) Recognize(image string, galleryName string) error {
	// TODO
	return nil
}

// Detect image attributes
func (k Kairos) Detect(image string) error {
	// TODO
	return nil
}

// List galleries
func (k Kairos) ListGalleries() error {
	// TODO
	return nil
}

// View a gallery
func (k Kairos) ViewGallery(galleryName string) error {
	// TODO
	return nil
}

// Remove a gallery
func (k Kairos) RemoveGallery(galleryName string) error {
	// TODO
	return nil
}

// Remove a subject
func (k Kairos) RemoveSubject(subjectId, galleryName string) error {
	// TODO
	return nil
}

func (k *Kairos) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	uri := k.BaseUrl.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set up header
	req.Header.Add("User-Agent", k.UserAgent)
	if appInfo := k.AppInfo; appInfo != nil {
		key, _ := appInfo.Password()
		req.Header.Add("app_id", appInfo.Username())
		req.Header.Add("app_key", key)
	}

	return req, nil
}

func (k *Kairos) do(req *http.Request, v interface{}) error {
	resp, err := k.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	var errorR interface{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorR)
	}
	return fmt.Errorf("%+v", errorR)
}

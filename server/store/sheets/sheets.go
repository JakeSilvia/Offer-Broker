package sheets

import (
	"backup/server/settings"
	"backup/server/store"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	SheetID = "1I9ITqjow4i7OfPV9Ikmmm1GHu-lc2mgE5Z8NQDcY0IU"
)

type SheetResp struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"major_dimension"`
	Values         [][]string `json:"values"`
}

type Sheet struct {
	SheetID    string
	Sheet      string
	StartRange string
	StopRange  string
	Client     *http.Client
}

func NewSheet() (store.Store, error) {
	config := &jwt.Config{
		Email:        settings.CLIENT_EMAIL,
		PrivateKeyID: settings.PRIVATE_KEY_ID,
		PrivateKey:   []byte(settings.PRIVATE_KEY),
		TokenURL:     google.JWTTokenURL,
		Scopes:       []string{"https://www.googleapis.com/auth/spreadsheets"},
	}
	client := config.Client(context.Background())

	return &Sheet{
		SheetID: SheetID,
		Client:  client,
	}, nil
}

func (s *Sheet) doAppendRequest(payload interface{}) error {
	requestUrl, err := url.Parse(s.buildUrl(true))
	if err != nil {
		return err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	request := &http.Request{
		Method: http.MethodPost,
		URL:    requestUrl,
		Body:   ioutil.NopCloser(bytes.NewReader(body)),
	}

	_, err = s.makeRequest(request)
	return err
}

func (s *Sheet) doReadRequest() (*SheetResp, error) {
	requestUrl, err := url.Parse(s.buildUrl(false))
	if err != nil {
		return nil, err
	}

	request := &http.Request{
		URL:    requestUrl,
		Method: http.MethodGet,
	}
	return s.makeRequest(request)
}

func (s *Sheet) makeRequest(r *http.Request) (*SheetResp, error) {
	resp, err := s.Client.Do(r)
	if err != nil {
		return nil, err
	}

	bts, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK || err != nil {
		return nil, fmt.Errorf("error making request: %v, %v", err, string(bts))
	}

	sheetResp := &SheetResp{}
	err = json.Unmarshal(bts, sheetResp)
	if err != nil {
		return nil, err
	}

	return sheetResp, nil
}

func (s *Sheet) buildUrl(write bool) string {
	u := fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%v/values/%v!%v",
		s.SheetID,
		s.Sheet,
		s.StartRange,
	)
	if s.StopRange != "" {
		u += fmt.Sprintf(":%v", s.StopRange)
	}

	u += "?key=" + settings.API_KEY
	if write {
		u = u + "&valueInputOption=USER_ENTERED"
	}
	return u
}

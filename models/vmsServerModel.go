package models

import (
	"bytes"
	"encoding/json"
	"errors"
	logv "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type VmsServerModel struct{}

type VmsLoginBody struct {
	AccountID string `json:"accountID"`
	Password  string `json:"password"`
}

func (m *VmsServerModel) ConnectionTest(
	account string, pwd string, protocol string, host string) (err error) {
	resp, err := http.Get(protocol + "://" + host + "/ping")
	if err != nil {
		logv.Error(err)
		return err
		// handle error
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logv.Error(err)
		return err
		// handle error
	}

	//logv.Println(string(body))

	vmsLoginResponse := &vmsLoginResponse{}

	shortenData := VmsLoginBody{
		account,
		pwd,
	}
	ba, _ := json.Marshal(shortenData)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", protocol+"://"+host+"/api/v1/user/loginUser", bytes.NewBuffer(ba))
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	content, err := ioutil.ReadAll(res.Body)
	respBody := string(content)
	//fmt.Printf("Post request with json result: %s\n", respBody)
	errq := json.Unmarshal([]byte(respBody), vmsLoginResponse)
	_ = errq

	defer res.Body.Close()
	if vmsLoginResponse.Code != 0 {
		return errors.New(vmsLoginResponse.Message)
	}

	return err
}

type vmsLoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	User    string `json:"user"`
}

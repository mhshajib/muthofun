package muthofun

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//	Real implementation of Sms interface
type MuthofunConnection struct {
	Url string
}

/*
Set sms gateway username and password
*/
func (m *MuthofunConnection) Init(username string, password string) {
	m.Url = fmt.Sprintf("http://developer.muthofun.com/sms.php?username=%s&password=%s", parseSpecials(username), parseSpecials(password))
}

//Filter # and & in param value
func parseSpecials(raw string) string {
	response := ""
	for _, c := range raw {
		char := string(c)
		if char == "#" {
			response += "%23"
		} else if char == "&" {
			response += "%26"
		} else {
			response += char
		}
	}
	return response
}

//Send sms
func (m *MuthofunConnection) Send(mobiles []string, sms string, unicode string) (map[string]interface{}, error) {
	var result map[string]interface{}

	//Checking is there any mobile number
	if len(mobiles) < 1 {
		err := errors.New("Need to enter at least one mobile number")
		return result, err
	}

	//Sequencing mobile number in a comma separated value from array
	sequenced_mobile := ""
	for i := 0; i < len(mobiles); i++ {
		if sequenced_mobile == "" {
			sequenced_mobile += mobiles[i]
		} else {
			sequenced_mobile += fmt.Sprintf(",%s", mobiles[i])
		}
	}

	//Completed url for request
	generatedUrl := fmt.Sprintf("%s&mobiles=%s&sms=%s&uniccode=%s", m.Url, sequenced_mobile, url.QueryEscape(sms), unicode)
	//Making an http get request
	resp, err := http.Get(generatedUrl)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	//Reading response body
	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	//Returning result
	return result, nil
}

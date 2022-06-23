package vrchat

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	API_KEY                  = "JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26"
	STATUS_DESCRIPTION_LIMIT = 31
)

type VRChat struct {
	cookie                 string
	twoFactorAuthCookie    string
	isTwoFactorAuthEnabled bool
	user                   *UserData
}

func New() *VRChat {
	return &VRChat{}
}

func (v *VRChat) Login(username, password string) error {
	if username == "" || password == "" {
		return errors.New("missing username or password")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://vrchat.com/api/1/auth/user?apiKey=%s", API_KEY), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > http.StatusAccepted {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return errors.New("Invalid Username or Password")
		default:
			return errors.New("Unexpected Error Occurred")
		}
	}
	cookie := resp.Header.Get("set-cookie")
	if cookie == "" {
		return errors.New("missing cookie")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	v.SetCookie(cookie)
	var twoFactorData *UserLoginRequireTwoFactorAuthResponseBody
	err = json.Unmarshal(body, &twoFactorData)
	if err != nil {
		return err
	}
	// If 2Factor auth is required, it will match this interface
	if len(twoFactorData.RequiresTwoFactorAuth) != 0 {
		v.SetToFactorAuthEnabled(true)
	} else {
		var userData *UserData
		err = json.Unmarshal(body, &userData)
		if err != nil {
			return err
		}
		if userData.ID == "" {
			return errors.New("Unexpected error occurred")
		} else {
			v.user = userData
		}
	}
	return nil
}

func (v *VRChat) TwoFactorAuthenticate(code string) error {
	if code == "" {
		return errors.New("Missing Code")
	}
	if !v.IsLoggedIn() {
		return errors.New("Please login before trying to authenticate")
	}
	client := &http.Client{}
	payload := fmt.Sprintf(`{"code":"%v"}`, code)
	data := strings.NewReader(payload)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://vrchat.com/api/1/auth/twofactorauth/totp/verify?apiKey=%v", API_KEY), data)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	cookies := v.GetCookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > http.StatusAccepted {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			return errors.New("Invalid Verification Code")
		default:
			return errors.New("Unexpected Error Occurred")
		}
	}
	cookie := resp.Header.Get("set-cookie")
	if cookie == "" {
		return errors.New("missing cookie")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	v.SetTwoFactorAuthCookie(cookie)
	var res *TwoFactorAuthResponseBody
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	if !res.Verified {
		return errors.New("Unable to verify")
	}
	return v.GetUser()
}

func (v *VRChat) GetUser() error {
	if !v.IsLoggedIn() {
		return errors.New("Cannot fetch user if not logged in")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://vrchat.com/api/1/auth/user?apiKey=%s", API_KEY), nil)
	if err != nil {
		return err
	}
	cookies := v.GetCookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > http.StatusAccepted {
		return errors.New("Unexpected Error Occurred")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var userData *UserData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return err
	}
	if userData.ID == "" {
		return errors.New("Unexpected error occurred")
	} else {
		v.user = userData
	}
	return nil
}

func (v *VRChat) SetStatus(status string) error {
	if status == "" {
		return errors.New("Missing Status")
	}
	if len(status) > STATUS_DESCRIPTION_LIMIT {
		status = status[:STATUS_DESCRIPTION_LIMIT]
	}
	if !v.IsLoggedIn() {
		return errors.New("Please login before to set status")
	}
	client := &http.Client{}
	payload := fmt.Sprintf(`{"statusDescription":"%v"}`, status)
	data := strings.NewReader(payload)

	req, err := http.NewRequest("PUT", fmt.Sprintf("https://vrchat.com/api/1/users/%v?apiKey=%v", v.user.ID, API_KEY), data)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Content-Type", "application/json")
	cookies := v.GetCookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode > http.StatusAccepted {
		return errors.New("Unexpected Error Occurred")
	}
	return nil
}

func (v *VRChat) IsLoggedIn() bool {
	return v.cookie != ""
}

func (v *VRChat) IsTwoFactorAuthEnabled() bool {
	return v.isTwoFactorAuthEnabled
}

func (v *VRChat) SetToFactorAuthEnabled(isEnabled bool) {
	v.isTwoFactorAuthEnabled = isEnabled
}

func (v *VRChat) SetCookie(cookie string) {
	v.cookie = cookie
}
func (v *VRChat) SetTwoFactorAuthCookie(cookie string) {
	v.twoFactorAuthCookie = cookie
}

func (v *VRChat) Clear() {
	v.cookie = ""
	v.isTwoFactorAuthEnabled = false
}

func (v *VRChat) GetCookies() []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", fmt.Sprintf("apiKey=%v;", API_KEY))
	header.Add("Cookie", v.cookie)
	header.Add("Cookie", v.twoFactorAuthCookie)
	req := http.Request{Header: header}
	return req.Cookies()
}

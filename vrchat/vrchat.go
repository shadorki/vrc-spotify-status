package vrchat

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_KEY = "JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26"

type VRChat struct {
	cookie                 string
	isTwoFactorAuthEnabled bool
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
	var data *UserLoginResponseBody
	err = json.Unmarshal(body, &data)
	fmt.Println(string(body))
	if err != nil {
		return err
	}
	v.cookie = cookie
	v.isTwoFactorAuthEnabled = data.TwoFactorAuthEnabled

	return nil
}

func (v *VRChat) IsLoggedIn() bool {
	return v.cookie != ""
}

func (v *VRChat) IsTwoFactorAuthEnabled() bool {
	return v.isTwoFactorAuthEnabled
}

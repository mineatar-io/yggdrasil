package yggdrasil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UsernameLookup struct {
	Username string `json:"name"`
	UUID     string `json:"id"`
}

type UsernameHistory struct {
	Username    string `json:"name"`
	ChangedToAt int64  `json:"changedToAt,omitempty"`
}

func UsernameToUUID(username string) (*UsernameLookup, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 204 {
			return nil, nil
		}

		return nil, fmt.Errorf("yggdrasil: unexpected response: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &UsernameLookup{}

	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

func UsernamesToUUIDs(usernames []string) ([]UsernameLookup, error) {
	postBody, err := json.Marshal(usernames)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://api.mojang.com/profiles/minecraft", "application/json", bytes.NewReader(postBody))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("yggdrasil: unexpected response: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := make([]UsernameLookup, 0)

	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func GetUsernameHistory(uuid string) ([]UsernameHistory, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.mojang.com/user/profiles/%s/names", uuid))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 204 {
			return nil, nil
		}

		return nil, fmt.Errorf("yggdrasil: unexpected response: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := make([]UsernameHistory, 0)

	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}

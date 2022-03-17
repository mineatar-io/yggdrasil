package yggdrasil

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProfileTextures struct {
	UUID       string `json:"id"`
	Username   string `json:"name"`
	Legacy     bool   `json:"legacy"`
	Properties []struct {
		Name      string `json:"name"`
		Value     string `json:"value"`
		Signature string `json:"signature,omitempty"`
	} `json:"properties"`
}

type DecodedTextures struct {
	Timestamp         int64  `json:"timestamp"`
	UUID              string `json:"uuid"`
	Username          string `json:"username"`
	SignatureRequired bool   `json:"signatureRequired"`
	Textures          struct {
		Skin struct {
			URL      string `json:"url"`
			Metadata struct {
				Model string `json:"model"`
			} `json:"metadata,omitempty"`
		} `json:"SKIN,omitempty"`
		Cape struct {
			URL string `json:"url"`
		} `json:"CAPE,omitempty"`
	} `json:"textures"`
}

func GetProfileTextures(uuid string) (*ProfileTextures, error) {
	resp, err := http.Get(fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid))

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

	response := &ProfileTextures{}

	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

func GetDecodedTexturesValue(value string) (*DecodedTextures, error) {
	decodedResult, err := base64.StdEncoding.DecodeString(value)

	if err != nil {
		return nil, err
	}

	result := &DecodedTextures{}

	err = json.Unmarshal(decodedResult, result)

	return result, err
}

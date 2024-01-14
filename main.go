package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/manifoldco/promptui"
)

type License struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Spdx_id string `json:"spdx_id"`
	Url     string `json:"url"`
	Node_id string `json:"node_id"`
}

func main() {
	licenses := getLicenses()
	parsedLicenses := parseLicenses(licenses)
	key := selector(parsedLicenses)
	license := getLicenseByKey(parsedLicenses, key)
	fmt.Println(license)
}

func getLicenses() []byte {
	res, err := http.Get("https://api.github.com/licenses")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func parseLicenses(licenses []byte) []License {
	var licenseList []License
	err := json.Unmarshal(licenses, &licenseList)
	if err != nil {
		fmt.Println(err)
	}
	return licenseList
}

func selector(parsedLicenses []License) string {
	var licenses []string
	for _, license := range parsedLicenses {
		licenses = append(licenses, license.Spdx_id)
	}

	prompt := promptui.Select{
		Label: "Choose your license:",
		Items: licenses,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}

func getLicenseByKey(parsedLicenses []License, key string) License {
	for _, license := range parsedLicenses {
		if license.Key == key {
			return license
		}
	}
	return License{}
}

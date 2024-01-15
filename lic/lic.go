package lic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type License struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Spdx_id string `json:"spdx_id"`
	Url     string `json:"url"`
	Node_id string `json:"node_id"`
}

type LicenseInfo struct {
	Key            string   `json:"key"`
	Name           string   `json:"name"`
	Spdx_id        string   `json:"spdx_id"`
	Url            string   `json:"url"`
	Node_id        string   `json:"node_id"`
	Html_url       string   `json:"html_url"`
	Description    string   `json:"description"`
	Implementation string   `json:"implementation"`
	Permissions    []string `json:"permissions"`
	Conditions     []string `json:"conditions"`
	Limitations    []string `json:"limitations"`
	Body           string   `json:"body"`
	Featured       bool     `json:"featured"`
}

func GetLicenses() []License {
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
	var licenseList []License
	errr := json.Unmarshal(body, &licenseList)
	if errr != nil {
		fmt.Println(err)
	}
	return licenseList
}

func GetLicenseByKey(parsedLicenses []License, key string) License {
	var license License
	for _, l := range parsedLicenses {
		if l.Name == key {
			license = l
			break
		}
	}
	return license
}

func GetLicenseInfo(license License) LicenseInfo {
	res, err := http.Get("https://api.github.com/licenses/" + license.Key)
	if err != nil {
		fmt.Println(err)
		return LicenseInfo{}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return LicenseInfo{}
	}
	var licenseInfo LicenseInfo
	errr := json.Unmarshal(body, &licenseInfo)
	if errr != nil {
		fmt.Println(err)
		return LicenseInfo{}
	}
	return licenseInfo
}

func MakeLicense(license LicenseInfo) {
	body := license.Body
	if license.Key == "mit" || license.Key == "bsd-2-clause" || license.Key == "bsd-3-clause" {
		name, year := getUserDetails()
		body = strings.ReplaceAll(body, "[fullname]", name)
		body = strings.ReplaceAll(body, "[year]", year)
	}
	err := os.WriteFile("LICENSE", []byte(body), 0644)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		os.Exit(1)
	}
	fmt.Println("License created successfully!")
}

func getUserDetails() (string, string) {
	nameReader := bufio.NewReader(os.Stdin)
	fmt.Print("Full name: ")
	name, _ := nameReader.ReadString('\n')
	yearReader := bufio.NewReader(os.Stdin)
	fmt.Print("Year: ")
	year, _ := yearReader.ReadString('\n')
	return strings.TrimSuffix(name, "\n"), strings.TrimSuffix(year, "\n")
}

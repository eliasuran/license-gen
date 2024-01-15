package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/mpu69/license-gen/lic"
)

func main() {
	licenses := lic.GetLicenses()

	key := selector(licenses)

	license := lic.GetLicenseByKey(licenses, key)

	licenseInfo := lic.GetLicenseInfo(license)

	lic.MakeLicense(licenseInfo)
}

func selector(parsedLicenses []lic.License) string {
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

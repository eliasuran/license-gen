package tests

import (
	"fmt"
	"testing"

	"github.com/eliasuran/license-generator/lic"
	"github.com/stretchr/testify/assert"
)

func TestGetLicences(t *testing.T) {
	t.Run("should return a slice with an unknown amount of structs with type of License", func(t *testing.T) {
		got := lic.GetLicenses()

		if got == nil {
			t.Log("The API returned a nil slice")
			fmt.Println("The API returned a nil slice")
			return
		}

		assert.IsType(t, []lic.License{}, got, "The response should be a slice of License")
		for _, license := range got {
			assert.IsType(t, lic.License{}, license, "Each element should be of type License")
		}
	})
}

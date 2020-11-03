package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

func getNotSupportedProvinces() []string {
	return []string{"PR", "SC", "RS"}
}

//WarrantyProvinceValidation is responsible for validation province warranty
func WarrantyProvinceValidation(province string) infra.Validation {
	notSupportedProvinces := getNotSupportedProvinces()
	for _, notSupportedProvince := range notSupportedProvinces {
		if province == notSupportedProvince {
			return infra.Validation{Err: fmt.Errorf("warranty province %s not supported", province)}
		}
	}
	return infra.Validation{Err: nil}
}

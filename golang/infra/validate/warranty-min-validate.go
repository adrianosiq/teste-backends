package validate

import (
	"fmt"

	"github.com/adrianosiq/teste-backends/golang/infra"
)

const (
	minWarrantyNumber = 1
)

//WarrantyMinValidation is responsible for validating the minimum number of warranties
func WarrantyMinValidation(number int) infra.Validation {
	if number < minWarrantyNumber {
		return infra.Validation{Err: fmt.Errorf("warranties number is less than minimum of %d", minWarrantyNumber)}
	}
	return infra.Validation{Err: nil}
}

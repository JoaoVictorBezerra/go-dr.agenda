package customErrors

import "errors"

var InsuranceNotFoundError = errors.New("insurance not found")
var UpdateInsuranceError = errors.New("update insurance error")

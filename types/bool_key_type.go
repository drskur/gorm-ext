package types

import (
	"strings"
	"github.com/pkg/errors"
)

type BoolKey string

func (bk BoolKey) Parse() (bool, error) {
	if lower := strings.ToLower(string(bk)); lower == "true" {
		return true, nil
	} else if lower == "false" {
		return false, nil
	} else {
		return false, errors.Errorf("cannot parse %s to bool", bk)
	}
}